import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { MatTableModule } from '@angular/material/table';
import { MatIconModule } from '@angular/material/icon';
import { TreeTableComponent } from './component/tree-table.component';
export { Node, Options } from './models';

@NgModule({
  declarations: [
    TreeTableComponent
  ],
  imports: [
    CommonModule,
    MatTableModule,
    MatIconModule
  ],
  exports: [
    TreeTableComponent
  ]
})
export class TreeTableModule { }
