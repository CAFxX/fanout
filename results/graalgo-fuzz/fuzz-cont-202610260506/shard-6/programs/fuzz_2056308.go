package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32 }
func genFn0() int {
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := []int{(fi8 % (((fi7 >> ((991) & 63))) | 1)), fi7, (-(fi8)), 32768}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if false {
        var v10 uint32 = 2147483647
        v11 := GenS1{F0: -731, F1: 1, F2: -333}
        _ = v10
        _ = v11
    }
    var v12 float64 = (-2.25 / 3.14159)
    v13 := GenS1{F0: 0, F1: 0, F2: 596}
    if false {
        var v14 float32 = (1.5 * 1e10)
        _ = v14
    } else {
        for fi15 := 0; fi15 < 16; fi15++ {
            fmt.Println(v12, v12)
            var v16 uint32 = 2147483648
            _ = fi15
            _ = v16
        }
        var v17 float64 = (v12 + ((0.1 * 0.1) - 1e10))
        _ = v17
    }
    _ = v12
    _ = v13
    return 269
}
func genFn1(p0 uint) int32 {
    if ((42 != -9223372036854775808) && ((648 > genFn0()) || false)) {
        for fi18 := 0; fi18 < 18; fi18++ {
            p0 -= 4294967295
            _ = fi18
        }
        fmt.Println(p0)
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        if (("foo bar" + "") <= "foo bar") {
            p0 = 106
        } else {
            var v20 float64 = -2.25
            _ = v20
        }
        var v21 uint8 = 185
        _ = fi19
        _ = v21
    }
    _ = p0
    return -90
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ("\n\t" + "foo bar")
    v3 := []string{((v2 + v2) + "hello"), "a", v2}
    var v4 uint8 = 48
    var v5 uint32 = 674
    var v6 uint = 0
    v22 := GenS2{F0: v2}
    v6 = 42
    if len(v3) > 0 {
        v3[0] = (v2 + ((v2 + "\n\t") + v2))
    }
    v23 := GenS3{F0: 1000}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v24 uint8 = v4
    var v25 uint16 = 435
    v26 := GenS1{F0: 0, F1: 377, F2: -740}
    if false {
        for fi27 := 0; fi27 < 10; fi27++ {
            v5 += (v5 / ((46) | 1))
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 bool = false
                _ = fi28
                _ = v29
            }
            var v30 uint16 = (^(v25))
            _ = fi27
            _ = v30
        }
        if (2147483647 <= 914) {
            switch -1 {
            case 42:
                v24 = 42
            case 32768:
                v4 -= 2
            }
            fmt.Println(v6, v6)
            for fi31 := 0; fi31 < 4; fi31++ {
                v32 := GenS4{F0: 751}
                v33 := GenS2{F0: (v2 + v2)}
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        if false {
            var v34 string = "\n\t"
            var v35 float64 = ((-2.25 / 1.0) + float64(1e10))
            _ = v34
            _ = v35
        }
    } else {
        var v36 int32 = -2147483648
        _ = v36
    }
    var v37 int8 = 42
    v24 = v24
    v2 = (v2 + "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
