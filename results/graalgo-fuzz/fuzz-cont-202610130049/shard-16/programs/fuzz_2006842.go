package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int, p1 int16, p2 int8) uint8 {
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint8 = 128
            var v10 string = ("hello" + "")
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS1{F0: false, F1: "foo bar", F2: 127}
            p0 = -2147483648
            switch "foo bar" {
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    switch (("foo bar" + ("a" + "hello")) + "\n\t") {
                    case "hello":
                        var v13 int8 = 100
                        _ = v13
                    default:
                        if false {
                            var v14 int32 = 42
                            _ = v14
                        }
                    }
                case "\n\t":
                    v15 := []int{fi11, (fi11 - fi7)}
                    _ = v15
                }
            }
            _ = fi11
            _ = v12
        }
        p2 = 42
        _ = fi7
    }
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint = 255
    v2 := []uint32{65535, 22}
    var v3 bool = true
    var v4 float32 = 3.14159
    v5 := []uint32{2, 0, 42, 4294967295}
    v6 := make(map[int]int)
    var v16 int16 = -1
    switch -2147483648 {
    case 32768:
        v3 = v3
    case -9223372036854775808:
        v16 = 0
    default:
        var v17 uint32 = 508
        _ = v17
    }
    fmt.Println(v4)
    switch 4294967295 {
    case -730:
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 int32 = 221
            if len(v2) > 0 {
                v2[0] = 1
            }
            _ = fi18
            _ = v19
        }
    default:
        v20 := GenS2{F0: 42}
        _ = v20
    }
    var v21 uint = 42
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(v21, v3, fi22)
        v4 -= v4
        for fi23 := 0; fi23 < 4; fi23++ {
            v3 = v3
            var v24 int32 = 42
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    v25 := GenS1{F0: v3, F1: ("\n\t" + (("foo bar" + "hello") + "")), F2: 15}
    v26 := GenS2{F0: 15}
    var v27 uint16 = 0
    var v28 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
