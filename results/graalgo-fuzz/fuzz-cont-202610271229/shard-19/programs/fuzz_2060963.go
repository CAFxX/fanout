package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
func genFn0(p0 int, p1 uint, p2 uint) int16 {
    v4 := []int{-1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    p0 = (p0 * 42)
    v5 := GenS1{F0: 110}
    var v6 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint64, p1 string, p2 uint64) bool {
    if (((0.1 * 0.0) / -2.25) > 0.1) {
        v7 := []int{-9223372036854775808}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS2{F0: ("" + p1)}
            _ = fi8
            _ = v9
        }
    }
    if (101 != 100) {
        var v10 float32 = 1e10
        _ = v10
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS1{F0: -86}
            v13 := []int{}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return ((42 == 255) && (!false))
}

func main() {
    var v1 uint = 0
    var v2 string = ""
    var v3 int32 = -1
    var v14 uint16 = 32767
    fmt.Println(v2, v1)
    if (!(0.0 > (1e10 - 3.14159))) {
        var v15 int8 = -1
        v2 = v2
        _ = v15
    }
    v16 := GenS1{F0: 127}
    fmt.Println(v3, v2, v2)
    v1 -= v1
    var v17 uint64 = 0
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 uint64 = (v17 * v17)
            var v21 uint64 = v17
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    if true {
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 string = "\n\t"
            _ = fi22
            _ = v23
        }
        switch 32768 {
        case -965:
            fmt.Println(v1, v14)
        case 2147483647:
            fmt.Println(v14, v3, v17)
        case 127:
            v24 := GenS2{F0: ((v2 + ("\n\t" + v2)) + (((v2 + "\n\t") + (v2 + "a")) + v2))}
            _ = v24
        default:
            for fi25 := 0; fi25 < 15; fi25++ {
                for fi26 := 0; fi26 < 16; fi26++ {
                    if false {
                        v27 := GenS1{F0: -128}
                        switch 128 {
                        case 0:
                            fmt.Println(v17, fi25, v3)
                        }
                        _ = v27
                    } else {
                        v28 := []int{fi26, (fi25 & (fi25 | (fi25 / ((fi25) | 1)))), fi25, 935}
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                        _ = v28
                    }
                    fmt.Println(v17, v2)
                    _ = fi26
                }
                _ = fi25
            }
        }
        switch -9223372036854775808 {
        case 976:
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 uint64 = (^(v17))
                _ = fi29
                _ = v30
            }
        case -2147483648:
            v31 := []int{}
            _ = v31
        }
    }
    v32 := GenS2{F0: "a"}
    switch v2 {
    case "foo bar":
        var v33 uint8 = 128
        _ = v33
    }
    v34 := []int{840}
    v35 := GenS2{F0: (v2 + "\n\t")}
    if len(v34) > 0 {
        v34[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
}
