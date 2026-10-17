package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint32, p1 uint64, p2 int32) int32 {
    switch ("" + "hello") {
    case "foo bar":
        var v7 uint = uint(p0)
        _ = v7
    default:
        switch (("\n\t" + "a") + ("" + ("a" + "hello"))) {
        case "hello":
            if false {
                v8 := GenS1{F0: -9223372036854775808}
                _ = v8
            } else {
                var v9 bool = true
                v9 = v9
                v10 := []int{128, 32767, -2147483648}
                _ = v10
                _ = v9
            }
        case "a":
            switch 979 {
            case -628:
                switch 0 {
                case 128:
                    for fi11 := 0; fi11 < 1; fi11++ {
                        p2 = 2147483647
                        var v12 int16 = -1
                        _ = fi11
                        _ = v12
                    }
                }
            }
        case "\n\t":
            fmt.Println(p2, p2)
        }
    }
    var v13 bool = (false || false)
    fmt.Println(v13, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    v3 := []uint32{394}
    var v4 uint8 = 42
    var v5 float32 = 1e10
    v6 := make(map[string]int)
    v5 += 1.5
    for fi14 := 0; fi14 < 15; fi14++ {
        if (false && (!true)) {
            for fi15 := 0; fi15 < 19; fi15++ {
                switch 32767 {
                case 32767:
                    v4 = 1
                case 42:
                    var v16 float32 = (0.1 * (v5 + v5))
                    if false {
                        var v17 uint = 1
                        var v18 int64 = -9223372036854775808
                        _ = v17
                        _ = v18
                    } else {
                        v5 = v16
                    }
                    v5 += 0.1
                    _ = v16
                case 1:
                    v5 += 1e10
                }
                _ = fi15
            }
            if true {
                if false {
                    v5 -= 1e10
                    v4 = (v4 % ((v4) | 1))
                    var v19 int16 = 1000
                    _ = v19
                }
                var v20 string = ("hello" + "a")
                _ = v20
            }
        } else {
            v21 := GenS1{F0: -1}
            switch fi14 {
            case 773:
                v5 -= (1e10 - v5)
            }
            _ = v21
        }
        v22 := GenS1{F0: -1}
        _ = fi14
        _ = v22
    }
    var v23 uint = 373
    for fi24 := 0; fi24 < 17; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            switch "foo bar" {
            case "a":
                v23 -= 255
            case "foo bar":
                v4 += (^(v4))
            case "":
                delete(v6, (("a" + "\n\t") + "hello"))
            default:
                v5 += (v5 * float32(float64((float32(1.5) - (v5 * 0.1)))))
            }
            _ = fi25
        }
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 float64 = 0.1
                v28 := GenS1{F0: -9223372036854775808}
                _ = fi26
                _ = v27
                _ = v28
            }
        default:
            if (723 > 565) {
                var v29 int32 = 42
                v30 := GenS2{F0: 18446744073709551615}
                _ = v29
                _ = v30
            }
        }
        _ = fi24
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        if (!(255 > 968)) {
            v23 -= (v23 + (^(v23)))
            _ = len(v3)
            if true {
                var v32 float64 = 1.5
                v5 += v5
                var v33 uint32 = 2147483647
                _ = v32
                _ = v33
            } else {
                v34 := GenS2{F0: 1}
                v5 += v5
                if false {
                    var v35 string = (((("a" + "a") + "hello") + ("a" + "a")) + "a")
                    v36 := GenS1{F0: 1}
                    var v37 uint16 = 255
                    _ = v35
                    _ = v36
                    _ = v37
                }
                _ = v34
            }
        }
        _ = fi31
    }
    v38 := GenS2{F0: 896}
    v5 = v5
    if false {
        var v39 uint16 = 32767
        fmt.Println(v39, v5)
        _ = v39
    }
    fmt.Println(v23, v23)
    for fi40 := 0; fi40 < 20; fi40++ {
        if true {
            v41 := GenS2{F0: 4294967295}
            var v42 int16 = 1000
            _ = v41
            _ = v42
        }
        if (255 >= 42) {
            fmt.Println(fi40)
            var v43 int32 = 2147483647
            _ = v43
        }
        v4 = (v4 ^ 58)
        _ = fi40
    }
    var v44 int64 = -857
    fmt.Println(v4, v5)
    v4 = (v4 * (^(v4)))
    var v45 float32 = (0.0 - (v5 * v5))
    fmt.Println(v4, v44)
    if true {
        v5 = (float32((3.14159 - (-2.25 - 3.14159))) + (v45 - v5))
    }
    var v46 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
