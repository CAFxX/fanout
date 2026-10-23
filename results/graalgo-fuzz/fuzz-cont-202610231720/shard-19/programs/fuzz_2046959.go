package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int64 = -1
    var v2 float32 = 3.14159
    var v3 uint8 = 227
    var v4 int = int(v3)
    if false {
        v5 := []int{(v4 >> ((182) & 63)), 287, v4}
        _ = v5
    }
    v2 += v2
    v1 = 42
    var v6 uint16 = 0
    var v7 uint32 = 65535
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := []int{fi8, (v4 & (fi8 % ((v4) | 1)))}
        _ = fi8
        _ = v9
    }
    if false {
        v10 := []int{v4}
        switch "\n\t" {
        case "foo bar":
            v10 = append(v10, v4)
        case "hello":
            fmt.Println(v2)
        }
        _ = v10
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            switch fi11 {
            case -909:
                v12 := GenS1{F0: 4294967295}
                _ = v12
            case -2147483648:
                v13 := GenS1{F0: 555}
                _ = v13
            default:
                switch ("a" + "foo bar") {
                case "":
                    if (true || (0 == -2)) {
                        var v14 uint16 = 276
                        var v15 uint32 = (v7 & (v7 >> ((42) & 63)))
                        v16 := []int{v4, (fi11 * v4)}
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        v17 := []int{0}
                        _ = v17
                    }
                case "a":
                    if (false && true) {
                        switch "hello" {
                        case "hello":
                            var v18 int = (v4 & (^(fi11)))
                            _ = v18
                        case "":
                            fmt.Println(v1, v4, fi11)
                        }
                    } else {
                        v2 -= v2
                        fmt.Println(v1, v6, v1)
                    }
                }
            }
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 uint64 = 42
                v21 := GenS1{F0: 1}
                v4 = v4
                _ = fi19
                _ = v20
                _ = v21
            }
            var v22 uint64 = uint64(v3)
            _ = fi11
            _ = v22
        }
    }
    var v23 uint32 = v7
    v24 := []int{0, v4, v4, v4}
    v25 := GenS1{F0: 1}
    v24 = append(v24, (v4 * (^(v4))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
