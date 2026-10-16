package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[int]int)
    var v2 bool = ((1 <= 42) && (3.14159 != 0.0))
    var v3 int = 1
    fmt.Println(v3)
    v2 = (42 != -266)
    switch "" {
    case "\n\t":
        for fi4 := 0; fi4 < 15; fi4++ {
            if (1 != 1) {
                var v5 int = fi4
                if false {
                    for fi6 := 0; fi6 < 10; fi6++ {
                        var v7 uint16 = 32767
                        v2 = (42 == 433)
                        fmt.Println(v7)
                        _ = fi6
                        _ = v7
                    }
                    var v8 uint64 = 0
                    _ = v8
                }
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 string = "\n\t"
                    v11 := GenS1{F0: 4294967295}
                    fmt.Println(v2, fi9, fi9)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = v5
            }
            for fi12 := 0; fi12 < 14; fi12++ {
                fmt.Println(v3)
                v13 := GenS2{F0: 2}
                _ = fi12
                _ = v13
            }
            var v14 uint16 = 1
            _ = fi4
            _ = v14
        }
    case "a":
        switch (v3 % (((v3 << ((0) & 63))) | 1)) {
        case 769:
            v2 = true
        case 127:
            v2 = (-1 <= 100)
        default:
            var v15 uint64 = 2
            _ = v15
        }
    }
    fmt.Println(v3, v3)
    v16 := GenS1{F0: 0}
    fmt.Println(v2, v3)
    var v17 bool = false
    v17 = (32767 >= 564)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
