package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint16, p1 int64) uint32 {
    v5 := GenS1{F0: -2147483648}
    if false {
        if false {
            v6 := []int{481, 2147483647}
            p1 = -263
            v7 := GenS1{F0: -1}
            _ = v6
            _ = v7
        }
        if false {
            fmt.Println(p1, p1)
            switch -9223372036854775808 {
            case 233:
                v8 := GenS2{F0: 42}
                _ = v8
            case 0:
                var v9 int8 = 42
                _ = v9
            case 9223372036854775807:
                v10 := GenS3{F0: p1}
                _ = v10
            default:
                fmt.Println(p0, p0)
            }
            p0 = (^(p0))
        } else {
            v11 := GenS1{F0: 251}
            for fi12 := 0; fi12 < 7; fi12++ {
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 uint16 = p0
                    if true {
                        v15 := GenS3{F0: (p1 | (p1 >> ((0) & 63)))}
                        p0 = v14
                        v14 = p0
                        _ = v15
                    }
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
            _ = v11
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint = 1
            fmt.Println(v17, p0)
            _ = fi16
            _ = v17
        }
    } else {
        v18 := GenS1{F0: 128}
        var v19 uint64 = 0
        var v20 uint8 = uint8(p0)
        _ = v18
        _ = v19
        _ = v20
    }
    fmt.Println(p0, p1)
    switch ("" + "a") {
    case "a":
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := []int{fi21}
            fmt.Println(p1, fi21, fi21)
            var v23 float32 = 1.5
            _ = fi21
            _ = v22
            _ = v23
        }
    case "":
        switch 32767 {
        case 0:
            var v24 float32 = 1.5
            _ = v24
        case 1:
            if false {
                var v25 uint = 4294967295
                v25 = 255
                p0 += 255
                _ = v25
            } else {
                var v26 int64 = (^(p1))
                _ = v26
            }
        case 4294967295:
            fmt.Println(p1, p0)
        }
    case "foo bar":
        var v27 int8 = -61
        _ = v27
    }
    v28 := []int{-2147483648, -1}
    _ = p0
    _ = p1
    _ = v28
    _ = v5
    return 65535
}

func main() {
    var v1 uint8 = 178
    v2 := make(map[string]int)
    var v3 int = 4294967295
    var v4 int64 = 1
    var v29 uint64 = 18446744073709551615
    v4 += v4
    v1 = (v1 >> ((55) & 63))
    v30 := GenS3{F0: -9223372036854775808}
    if (-1 <= -416) {
        if (1.5 >= (1e10 - (0.1 * (-2.25 - ((float64((3.14159 + 1e10)) - 1.5) / 0.1))))) {
            v31 := []int{-266}
            delete(v2, "")
            _ = v31
        } else {
            var v32 float64 = (1.5 * ((1.5 / 1.0) - 0.1))
            _ = v32
        }
    } else {
        v33 := []int{int(v4)}
        _ = v33
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        v35 := []int{-458}
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        v3 -= (fi36 << ((30) & 63))
        v37 := GenS4{F0: 126, F1: (100 < 107)}
        v38 := GenS4{F0: (^(v1)), F1: true}
        _ = fi36
        _ = v37
        _ = v38
    }
    if (false && (v29 >= (^(v29)))) {
        v39 := GenS5{F0: (v29 & v29), F1: uint8(v29)}
        for fi40 := 0; fi40 < 14; fi40++ {
            v29 -= v29
            _ = fi40
        }
        var v41 int16 = 430
        _ = v39
        _ = v41
    } else {
        if (!false) {
            var v42 uint = 357
            var v43 string = (("foo bar" + "hello") + "a")
            _ = v42
            _ = v43
        }
        switch "\n\t" {
        case "foo bar":
            var v44 int64 = (v4 * 794)
            _ = v44
        }
        var v45 uint8 = v1
        _ = v45
    }
    switch ("" + (("\n\t" + ("" + "foo bar")) + ("\n\t" + ""))) {
    case "a":
        fmt.Println(v1, v1, v29)
    case "hello":
        switch "hello" {
        case "a":
            var v46 string = "foo bar"
            _ = v46
        }
    default:
        switch "\n\t" {
        case "":
            if (!(v29 <= (v29 << ((uint(v29)) & 63)))) {
                var v47 int32 = int32(v4)
                v48 := GenS1{F0: (-(v3))}
                var v49 int64 = 1
                _ = v47
                _ = v48
                _ = v49
            }
        case "hello":
            v50 := []int{0, v3, (v3 ^ v3)}
            _ = v50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
