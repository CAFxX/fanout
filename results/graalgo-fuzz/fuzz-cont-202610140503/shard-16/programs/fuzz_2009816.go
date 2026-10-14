package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 0.0
    v3 := make(map[uint32]string)
    var v4 bool = ("\n\t" >= "")
    var v5 int64 = -643
    v6 := make(map[int]bool)
    v7 := GenS1{F0: (v5 % ((42) | 1))}
    fmt.Println(v2, v2)
    for fi8 := 0; fi8 < 9; fi8++ {
        if ((65535 < 782) || (0 > 2147483648)) {
            fmt.Println(v2, v4, fi8)
            v5 = 42
        } else {
            v2 = v2
            switch (^(fi8)) {
            case 1:
                fmt.Println(v2, v5, fi8)
            case 32768:
                if true {
                    if (0 < -1) {
                        for fi9 := 0; fi9 < 17; fi9++ {
                            _ = len(v1)
                            _ = fi9
                        }
                    }
                    var v10 uint16 = 32767
                    _ = v10
                }
            case -2147483648:
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 int = fi8
                    _ = fi11
                    _ = v12
                }
            default:
                delete(v1, 4294967295)
            }
        }
        _ = fi8
    }
    fmt.Println(v2, v4, v4)
    var v13 int32 = -2147483648
    v13 += (v13 * v13)
    v13 = v13
    if v4 {
        switch 127 {
        case 1:
            var v14 int16 = 42
            _ = v14
        case 2147483647:
            v2 = 0.1
        case 319:
            var v15 uint16 = 396
            _ = v15
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            fmt.Println(v4, v13, v5)
            switch "a" {
            case "foo bar":
                var v17 uint32 = 664
                _ = v17
            case "":
                v18 := GenS1{F0: v5}
                _ = v18
            case "hello":
                for fi19 := 0; fi19 < 2; fi19++ {
                    fmt.Println(v5, v4, v4)
                    _ = fi19
                }
            }
            _ = fi16
        }
        v20 := []int{}
        _ = v20
    } else {
        if true {
            var v21 uint64 = 78
            for fi22 := 0; fi22 < 8; fi22++ {
                v23 := []int{(^(fi22)), (fi22 & (fi22 & fi22))}
                fmt.Println(v4, fi22)
                v24 := GenS1{F0: v5}
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = v21
        }
        v13 -= 2147483647
        var v25 float64 = ((float64(-2.25) - 3.14159) - float64(1.5))
        _ = v25
    }
    v26 := GenS1{F0: 0}
    delete(v3, 2147483648)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
