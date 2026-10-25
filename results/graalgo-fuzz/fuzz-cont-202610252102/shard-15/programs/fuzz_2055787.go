package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0(p0 int64, p1 int64, p2 uint8) uint8 {
    var v7 string = (("foo bar" + "\n\t") + "foo bar")
    if false {
        p2 = (p2 + (p2 & p2))
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 string = ("foo bar" + ("foo bar" + "a"))
            var v10 uint32 = 4294967295
            p1 += (^(p1))
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        var v11 string = ""
        v12 := GenS1{F0: (!(1 >= 4294967295))}
        _ = v11
        _ = v12
    }
    v13 := []int{208, -2147483648}
    if false {
        p2 += 1
        if false {
            if (((3.14159 - 3.14159) / 1e10) == (0.0 - (3.14159 * float64(1e10)))) {
                p1 += p0
                var v14 int32 = 0
                v14 -= 0
                _ = v14
            }
        } else {
            var v15 bool = false
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            v16 := GenS1{F0: v15}
            _ = v15
            _ = v16
        }
        if (true && (true && (!false))) {
            v13 = append(v13, 0)
            p2 += (p2 >> ((65535) & 63))
            fmt.Println(p0)
        } else {
            p1 -= 9223372036854775807
            v17 := GenS1{F0: true}
            if len(v13) > 0 {
                v13[0] = 0
            }
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(p2, v7, fi18)
            _ = fi18
        }
        if (true && false) {
            var v19 string = "hello"
            var v20 int32 = 2147483647
            v20 = -2147483648
            _ = v19
            _ = v20
        }
    }
    v21 := GenS1{F0: (!(!(true || (p0 < (^(p0))))))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v21
    _ = v7
    return (^(p2))
}
func genFn1(p0 int16, p1 int) int32 {
    for fi22 := 0; fi22 < 12; fi22++ {
        for fi23 := 0; fi23 < 18; fi23++ {
            v24 := GenS1{F0: true}
            _ = fi23
            _ = v24
        }
        fmt.Println(fi22, p0, fi22)
        var v25 bool = true
        _ = fi22
        _ = v25
    }
    fmt.Println(p0, p0)
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 float32 = (1.5 * 1.5)
    var v2 string = ""
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    v5 := []uint32{756, 4294967295, 447}
    var v6 uint16 = 2
    v6 = v6
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    v6 -= v6
    _ = len(v4)
    for fi26 := 0; fi26 < 17; fi26++ {
        v2 = v2
        fmt.Println(v1, v1)
        _ = fi26
    }
    var v27 uint = 2
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 17; fi29++ {
            v3[457] = 4294967295
            _ = fi29
        }
        if true {
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
            v6 -= v6
        }
        var v30 uint16 = v6
        _ = fi28
        _ = v30
    }
    var v31 bool = (!(!true))
    v32 := GenS2{F0: v6, F1: -9223372036854775808, F2: 127}
    if true {
        var v33 int16 = 423
        var v34 string = "hello"
        _ = v33
        _ = v34
    }
    if false {
        v31 = v31
    }
    var v35 int = 42
    var v36 int = 386
    var v37 uint8 = uint8(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
