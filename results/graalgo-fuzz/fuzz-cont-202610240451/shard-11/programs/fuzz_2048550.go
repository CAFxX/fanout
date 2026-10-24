package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint16) uint16 {
    var v6 uint8 = 255
    p0 -= (p0 / ((p0) | 1))
    fmt.Println(v6, p0)
    _ = p0
    _ = v6
    return p0
}
func genFn1() uint16 {
    if false {
        var v7 bool = false
        if v7 {
            fmt.Println(v7)
            v7 = v7
            v7 = v7
        }
        _ = v7
    }
    if (!((-80 == 100) || (!(false && false)))) {
        var v8 int64 = -614
        v9 := []int{128, -9223372036854775808, 4294967295, 370}
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 bool = false
            switch -9223372036854775808 {
            case 32768:
                var v12 int8 = -125
                _ = v12
            case 127:
                v13 := GenS1{F0: 346}
                _ = v13
            }
            v14 := GenS1{F0: 10}
            _ = fi10
            _ = v11
            _ = v14
        }
        _ = v8
        _ = v9
    }
    v15 := GenS1{F0: 2}
    _ = v15
    return 0
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    var v5 int16 = -32768
    v5 -= v5
    var v16 int16 = v5
    v17 := v3[1]
    fmt.Println(v16)
    var v18 int8 = -1
    if true {
        switch "foo bar" {
        case "":
            for fi19 := 0; fi19 < 1; fi19++ {
                v2[(v17 | 227)] = (v1 >= float64((0.1 + -2.25)))
                v20 := GenS1{F0: 255}
                var v21 string = "\n\t"
                _ = fi19
                _ = v20
                _ = v21
            }
        case "a":
            var v22 uint8 = 255
            _ = v22
        default:
            fmt.Println(v18, v16)
        }
        switch ("a" + "") {
        case "":
            switch ((("" + "foo bar") + ("hello" + ("hello" + "\n\t"))) + ("foo bar" + "")) {
            case "":
                v5 += 32767
            default:
                v18 = 42
            }
        }
        switch 2147483647 {
        case -1:
            fmt.Println(v17, v16, v16)
        case -569:
            v23 := v4[2147483647]
            _ = v23
        }
    }
    if false {
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 float64 = ((float64(1e10) * 0.1) + (3.14159 + v1))
            switch "hello" {
            case "\n\t":
                v26 := []int{(fi24 << ((584) & 63)), fi24, (fi24 - (fi24 - fi24))}
                _ = v26
            default:
                v27 := GenS2{F0: 315}
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        switch 32767 {
        case 32:
            fmt.Println(v17)
        case 127:
            var v28 int8 = (v18 | (v18 << ((255) & 63)))
            _ = v28
        default:
            var v29 uint16 = genFn1()
            _ = v29
        }
    }
    var v30 int64 = -823
    fmt.Println(v18, v1)
    v30 += (v30 << ((42) & 63))
    var v31 int = int(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
