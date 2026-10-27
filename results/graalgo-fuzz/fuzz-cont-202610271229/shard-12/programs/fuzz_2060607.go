package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int; F1 int64 }
func genFn0(p0 int64, p1 uint16) int32 {
    v7 := GenS1{F0: 42, F1: 127}
    p0 = 1
    var v8 int32 = 719
    fmt.Println(p1)
    var v9 int32 = (-(v8))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 885
    var v3 string = ("a" + "hello")
    var v4 float64 = (3.14159 / 0.1)
    var v5 float32 = 0.0
    var v6 float64 = v4
    v10 := GenS2{F0: 42}
    switch ("a" + "\n\t") {
    case "hello":
        v4 = v4
    case "a":
        switch (v3 + ("\n\t" + v3)) {
        case "\n\t":
            v2 = 4294967295
        case "":
            if (0 == 42) {
                var v11 float64 = (((1e10 / 3.14159) + -2.25) - 1e10)
                var v12 uint64 = (v2 << ((241) & 63))
                _ = v11
                _ = v12
            } else {
                fmt.Println(v5)
                var v13 bool = (1 < 100)
                if (v6 >= -2.25) {
                    v14 := GenS2{F0: 2147483647}
                    fmt.Println(v13, v2)
                    if v13 {
                        var v15 uint64 = v2
                        switch -394 {
                        case -530:
                            var v16 string = "foo bar"
                            _ = v16
                        case 651:
                            v17 := GenS2{F0: 4294967295}
                            _ = v17
                        case 32768:
                            v18 := []int{4294967295}
                            _ = v18
                        }
                        _ = v15
                    }
                    _ = v14
                }
                _ = v13
            }
        default:
            var v19 float32 = float32(-2.25)
            _ = v19
        }
    case "foo bar":
        fmt.Println(v4, v5, v4)
    }
    switch v3 {
    case "\n\t":
        fmt.Println(v4, v4, v6)
    case "foo bar":
        v6 = v6
    }
    var v20 int = 0
    fmt.Println(v2, v2, v5)
    if (false && (!((v3 + v3) <= "a"))) {
        v5 -= v5
    } else {
        switch (v20 - int(v2)) {
        case 128:
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 float64 = v4
                _ = fi21
                _ = v22
            }
        case 472:
            switch (v20 << ((682) & 63)) {
            case 42:
                v23 := GenS2{F0: 2}
                _ = v23
            case 369:
                v20 -= v20
            default:
                var v24 float64 = float64(-2.25)
                _ = v24
            }
        }
        var v25 bool = false
        for fi26 := 0; fi26 < 12; fi26++ {
            fmt.Println(v2, v20)
            _ = fi26
        }
        _ = v25
    }
    v27 := v1[""]
    v20 = v20
    for fi28 := 0; fi28 < 1; fi28++ {
        v29 := v1[v3]
        var v30 uint16 = 255
        for fi31 := 0; fi31 < 5; fi31++ {
            switch v3 {
            case "hello":
                if (v2 != (^(v2))) {
                    v4 = ((v6 * (v6 * float64(v5))) / v4)
                    fmt.Println(v4, v2)
                } else {
                    var v32 string = v3
                    fmt.Println(fi28, fi31)
                    if false {
                        v33 := v1[v3]
                        v34 := GenS2{F0: 2147483648}
                        v35 := GenS3{F0: 1000, F1: 18446744073709551615}
                        _ = v33
                        _ = v34
                        _ = v35
                    } else {
                        switch (^(fi28)) {
                        case 9223372036854775807:
                            v29 = (!v29)
                            v36 := GenS2{F0: 0}
                            v30 -= v30
                            _ = v36
                        }
                    }
                    _ = v32
                }
            case "":
                v37 := GenS4{F0: 32767}
                _ = v37
            case "foo bar":
                var v38 uint16 = (v30 ^ v30)
                _ = v38
            default:
                v39 := GenS1{F0: 0, F1: (fi31 - (fi31 * v20))}
                _ = v39
            }
            v40 := GenS5{F0: 42, F1: 4294967295}
            var v41 float32 = ((v5 * v5) + ((float32(float64(v5)) / (float32((v6 + v4)) - (v5 / (0.1 * (1.5 - v5))))) * float32((v6 + v6))))
            _ = fi31
            _ = v40
            _ = v41
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println(v20)
    v42 := GenS6{F0: v20, F1: -9223372036854775808}
    v43 := GenS2{F0: 65535}
    for fi44 := 0; fi44 < 7; fi44++ {
        var v45 int8 = -31
        v46 := []int{}
        if v27 {
            var v47 uint8 = 107
            delete(v1, "\n\t")
            var v48 uint8 = v47
            _ = v47
            _ = v48
        } else {
            for fi49 := 0; fi49 < 19; fi49++ {
                v4 = float64(float32(1e10))
                v5 = (v5 + v5)
                v50 := GenS3{F0: -850, F1: v2}
                _ = fi49
                _ = v50
            }
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println(v6, v5)
    delete(v1, v3)
    for fi51 := 0; fi51 < 14; fi51++ {
        var v52 string = "hello"
        v2 = v2
        switch -787 {
        case 32768:
            v2 -= (^(v2))
        default:
            switch 9223372036854775807 {
            case 32768:
                v53 := []int{(v20 | (fi51 << ((4294967295) & 63)))}
                _ = v53
            case 42:
                v5 = 0.0
            default:
                var v54 int16 = -1
                _ = v54
            }
        }
        _ = fi51
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
