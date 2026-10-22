package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int) int32 {
    switch p0 {
    case 1:
        v5 := []int{p0, p0, p0}
        _ = v5
    default:
        var v6 uint = 2
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            switch "" {
            case "\n\t":
                var v9 uint16 = 42
                _ = v9
            }
            _ = fi8
        }
        var v10 bool = true
        v11 := GenS1{F0: 127}
        _ = fi7
        _ = v10
        _ = v11
    }
    _ = p0
    return 942
}

func main() {
    var v1 uint32 = 2
    var v2 uint8 = 104
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    if (65535 > uint16(v1)) {
        v1 = 2147483648
    } else {
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 int = fi12
            v1 = (v1 * (v1 - v1))
            v14 := GenS1{F0: 127}
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(v2)
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            if (2 == 255) {
                fmt.Println(v2, fi15, fi15)
            }
            v2 -= (v2 ^ (v2 * 0))
            v1 -= (^(v1))
            if true {
                delete(v3, ((("a" + ("" + "foo bar")) + "foo bar") + ("hello" + "")))
                v1 = (^(v1))
                v2 -= v2
            }
            _ = fi16
        }
        var v17 uint = 1
        var v18 uint32 = v1
        _ = fi15
        _ = v17
        _ = v18
    }
    fmt.Println(v1)
    var v19 float64 = 0.1
    v2 += 127
    if ((741 >= 255) && (false && (3.14159 <= v19))) {
        var v20 float64 = 0.1
        _ = v20
    } else {
        v21 := []int{-9223372036854775808, int(v1)}
        _ = len(v21)
        _ = v21
    }
    v22 := GenS1{F0: 1}
    v23 := GenS1{F0: 42}
    fmt.Println(v19, v2, v19)
    switch 127 {
    case 32768:
        v2 = (v2 >> ((907) & 63))
    }
    switch ("hello" + "hello") {
    case "a":
        delete(v3, "hello")
    }
    var v24 int = 9223372036854775807
    _ = len(v4)
    v2 = (v2 * v2)
    switch "\n\t" {
    case "a":
        fmt.Println(v2, v2, v2)
    default:
        if (-675 != -32768) {
            v1 = (^(v1))
            for fi25 := 0; fi25 < 18; fi25++ {
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 int8 = 100
                    v28 := GenS1{F0: 127}
                    v2 += uint8(v1)
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                if true {
                    v29 := GenS1{F0: 127}
                    var v30 uint8 = 255
                    if false {
                        v31 := GenS1{F0: -60}
                        _ = v31
                    } else {
                        var v32 int = v24
                        _ = v32
                    }
                    _ = v29
                    _ = v30
                } else {
                    fmt.Println(fi25, v19)
                    fmt.Println(v1, v2, v24)
                    var v33 int16 = 1
                    _ = v33
                }
                _ = fi25
            }
        }
    }
    v19 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
