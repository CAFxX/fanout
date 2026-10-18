package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int64, p1 int16, p2 bool) bool {
    v7 := GenS1{F0: -742}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 bool = p2
        switch fi8 {
        case -9223372036854775808:
            switch "hello" {
            case "foo bar":
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := GenS1{F0: int(p0)}
                    v12 := []int{fi10, (fi8 + 127), -744}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "a":
                for fi13 := 0; fi13 < 7; fi13++ {
                    p2 = p2
                    _ = fi13
                }
            case "hello":
                fmt.Println(v9, p2, p2)
            }
        case -228:
            v14 := GenS1{F0: (^(fi8))}
            _ = v14
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[int]bool)
    v3 := []int{4294967295}
    v4 := []int{42, v1, (^(v1))}
    var v5 string = "foo bar"
    var v6 int = v1
    switch (v6 & v6) {
    case -2147483648:
        v6 = v6
    case 32768:
        var v15 uint8 = 128
        _ = v15
    case -299:
        fmt.Println(v6, v5, v5)
    default:
        switch 699 {
        case 42:
            var v16 uint8 = 229
            _ = v16
        case -457:
            v17 := GenS2{F0: 566}
            _ = v17
        default:
            _ = len(v2)
        }
    }
    v6 = v1
    var v18 uint64 = 550
    var v19 uint16 = 128
    var v20 uint = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v21 := GenS1{F0: v1}
    if len(v4) > 0 {
        v4[0] = (v1 + v6)
    }
    if (-842 == 42) {
        v5 = (v5 + v5)
    } else {
        var v22 uint = (^(v20))
        v19 += uint16(v18)
        v6 -= v6
        _ = v22
    }
    var v23 uint64 = v18
    if false {
        v3 = append(v3, -1)
        for fi24 := 0; fi24 < 11; fi24++ {
            for fi25 := 0; fi25 < 17; fi25++ {
                v26 := v2[int(v20)]
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v27 int16 = 0
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi24
        }
    } else {
        switch ((v5 + "\n\t") + "\n\t") {
        case "foo bar":
            v1 -= 599
        }
        v2[v1] = false
        for fi28 := 0; fi28 < 20; fi28++ {
            fmt.Println(v19, v20, v19)
            _ = fi28
        }
    }
    fmt.Println(v23, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
