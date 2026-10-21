package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int; F2 uint16 }
type GenS4 struct { F0 int8; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 string }
func genFn0() string {
    var v4 int16 = 0
    var v5 int32 = -2147483648
    var v6 int = 9223372036854775807
    _ = v4
    _ = v5
    _ = v6
    return "foo bar"
}
func genFn1(p0 string) int {
    if false {
        switch p0 {
        case "foo bar":
            fmt.Println(p0)
        case "":
            p0 = ((p0 + "\n\t") + "\n\t")
        case "a":
            p0 = (((p0 + (("hello" + p0) + ((p0 + ("" + (p0 + (p0 + p0)))) + p0))) + ((("hello" + p0) + (("hello" + p0) + p0)) + (p0 + p0))) + (p0 + ("foo bar" + "foo bar")))
        }
    }
    v7 := []int{}
    v8 := GenS1{F0: 393, F1: 2}
    _ = p0
    _ = v7
    _ = v8
    return -654
}

func main() {
    v1 := []uint32{65535, 916}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v9 int32 = -3
    if ((!(0 <= 42)) && true) {
        fmt.Println(v9)
        if (!(!(!true))) {
            v10 := GenS2{F0: 2}
            for fi11 := 0; fi11 < 4; fi11++ {
                for fi12 := 0; fi12 < 12; fi12++ {
                    v13 := GenS2{F0: 128}
                    v14 := GenS3{F0: 2, F1: 2147483647, F2: 42}
                    fmt.Println(v9)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v15 := GenS4{F0: -113, F1: -1, F2: 0}
                v9 = 2147483647
                _ = fi11
                _ = v15
            }
            _ = v10
        }
    } else {
        if (!(!(42 != 42))) {
            switch "" {
            case "":
                fmt.Println(v9)
            }
            v9 += (v9 / ((v9) | 1))
        }
        delete(v3, "")
        var v16 int64 = -1
        _ = v16
    }
    var v17 uint = 0
    var v18 int8 = -128
    switch -155 {
    case 127:
        switch genFn1("") {
        case 42:
            if ("" != ("\n\t" + ("" + "\n\t"))) {
                v3[("hello" + "a")] = false
                var v19 int8 = v18
                _ = v19
            }
        case 467:
            switch 128 {
            case 32767:
                var v20 int8 = v18
                _ = v20
            }
        default:
            v21 := GenS5{F0: 42}
            _ = v21
        }
    }
    if true {
        v18 = v18
    }
    v17 = 18446744073709551615
    if false {
        for fi22 := 0; fi22 < 16; fi22++ {
            delete(v2, 0)
            v23 := GenS6{F0: v17}
            if ((!true) || false) {
                if false {
                    fmt.Println(v18)
                    var v24 uint64 = 42
                    _ = v24
                }
                fmt.Println(fi22, v9)
            } else {
                v18 = (^(v18))
                var v25 string = ""
                _ = v25
            }
            _ = fi22
            _ = v23
        }
    } else {
        fmt.Println(v18, v17)
        switch 82 {
        case 32767:
            fmt.Println(v17, v17)
        case 128:
            delete(v3, ("\n\t" + ""))
        case -1:
            fmt.Println(v17, v18, v17)
        default:
            switch 4294967295 {
            case 4294967295:
                var v26 int32 = (v9 % ((v9) | 1))
                _ = v26
            case 32768:
                var v27 float64 = 0.1
                _ = v27
            default:
                if ((!(true || true)) || false) {
                    var v28 float32 = float32(((3.14159 / float64(0.1)) / ((0.0 + float64(1e10)) - (1e10 * 1e10))))
                    v29 := GenS7{F0: ("" + "")}
                    v17 = 21
                    _ = v28
                    _ = v29
                }
            }
        }
        v18 = (v18 << (((^(v17))) & 63))
    }
    var v30 float64 = 1e10
    if (0.1 != float32((float64(float32(v30)) - ((1e10 - v30) + ((float64((1e10 + float32(v30))) * 0.0) + v30))))) {
        var v31 uint = (^(v17))
        var v32 int8 = v18
        var v33 uint64 = 1
        _ = v31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        if false {
            for fi35 := 0; fi35 < 7; fi35++ {
                v18 = (v18 ^ v18)
                v36 := GenS4{F0: (-(v18)), F1: 131, F2: 2147483648}
                var v37 uint8 = 92
                switch "foo bar" {
                case "hello":
                    delete(v2, 0)
                }
                _ = fi35
                _ = v36
                _ = v37
            }
            v17 += v17
        }
        switch ("hello" + "a") {
        case "":
            v38 := v2[2147483648]
            _ = v38
        case "hello":
            var v39 bool = true
            _ = v39
        case "\n\t":
            switch "a" {
            case "a":
                for fi40 := 0; fi40 < 20; fi40++ {
                    var v41 int64 = 9223372036854775807
                    _ = fi40
                    _ = v41
                }
            case "hello":
                v42 := GenS1{F0: -1, F1: 2}
                _ = v42
            case "\n\t":
                v30 = (v30 * 1.5)
            default:
                var v43 int64 = 0
                _ = v43
            }
        }
        _ = fi34
    }
    var v44 uint16 = 42
    var v45 float64 = 0.1
    if (2 >= 65) {
        v46 := GenS2{F0: 42}
        var v47 float32 = (((((3.14159 + -2.25) * float32((((v45 + v45) - ((v30 * 0.0) * v30)) + (v30 - v30)))) + float32(((float64(0.0) - v30) + 0.0))) * float32(((0.1 * float64((1.5 * 1.5))) - v45))) - (float32(0.1) + ((0.1 * -2.25) / 3.14159)))
        _ = v46
        _ = v47
    } else {
        delete(v3, ("a" + "hello"))
    }
    switch ((("" + "a") + (((("a" + ("hello" + "a")) + "hello") + ("hello" + "")) + "hello")) + "a") {
    case "a":
        v1 = append(v1, 1)
    case "hello":
        var v48 int = -53
        _ = v48
    default:
        for fi49 := 0; fi49 < 5; fi49++ {
            v50 := GenS3{F0: 49, F1: fi49, F2: 42}
            for fi51 := 0; fi51 < 10; fi51++ {
                v17 += v17
                var v52 int8 = -79
                _ = fi51
                _ = v52
            }
            _ = fi49
            _ = v50
        }
    }
    var v53 int8 = 100
    v18 += v18
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v53:", v53)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v44
    _ = v45
    _ = v53
    _ = v9
}
