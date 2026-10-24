package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }

func main() {
    v1 := make(map[int]int)
    var v2 int = 4294967295
    v3 := []int{128}
    if len(v3) > 0 {
        v3[0] = v2
    }
    v1[(v2 ^ (v2 % ((v2) | 1)))] = (-(v2))
    var v4 float32 = -2.25
    var v5 uint = 255
    if len(v3) > 0 {
        v3[0] = (v2 >> (((v5 % ((323) | 1))) & 63))
    }
    v4 = v4
    if true {
        var v6 uint32 = 680
        v6 = (v6 % (((v6 - 2)) | 1))
        _ = v6
    } else {
        var v7 int32 = 1
        if true {
            v8 := GenS1{F0: (v2 >> ((480) & 63)), F1: "a"}
            var v9 string = ("hello" + ("\n\t" + "a"))
            _ = len(v3)
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 bool = true
                _ = fi10
                _ = v11
            }
            _ = len(v1)
            for fi12 := 0; fi12 < 10; fi12++ {
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 int32 = 850
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
        }
        _ = v7
    }
    v15 := GenS1{F0: 86, F1: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
