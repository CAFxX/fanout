package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int16 = 0
    v2 := make(map[int]bool)
    v3 := []int{1, 1, 42}
    v4 := []int{2147483647, 127}
    v1 += (v1 * v1)
    var v5 int16 = (v1 - v1)
    if true {
        switch "a" {
        case "hello":
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v1 = (v5 / (((v1 ^ 32767)) | 1))
                v6 := GenS1{F0: (v1 * 32767)}
                _ = v6
            } else {
                for fi7 := 0; fi7 < 11; fi7++ {
                    var v8 string = (("\n\t" + ("a" + "hello")) + "foo bar")
                    v8 = "foo bar"
                    _ = fi7
                    _ = v8
                }
                var v9 int = 32768
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 uint8 = 127
            var v12 uint16 = 255
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 int64 = 324
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    }
    var v16 uint8 = 255
    fmt.Println(v1, v5)
    v5 -= (^(v5))
    switch 127 {
    case 42:
        var v17 uint8 = v16
        _ = v17
    case -46:
        v1 += -1
    case -782:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v18 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
