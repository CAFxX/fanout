package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32 }
func genFn0(p0 string, p1 int16) uint {
    v4 := GenS1{F0: 826}
    var v5 int = 365
    v6 := []int{32767, (v5 / (((v5 & 42)) | 1)), v5}
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS2{F0: 845}
        if false {
            v9 := GenS2{F0: 0}
            v10 := GenS3{F0: "\n\t"}
            var v11 int = fi7
            _ = v10
            _ = v11
            _ = v9
        }
        switch v5 {
        case -2147483648:
            fmt.Println(v5, fi7)
        default:
            var v12 int64 = -150
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 2
    v3 := []uint32{42}
    switch 74 {
    case 4294967295:
        switch "hello" {
        case "a":
            v2 -= (v2 | 42)
        case "\n\t":
            v13 := GenS3{F0: (("a" + "hello") + (("a" + (("\n\t" + "hello") + "")) + "foo bar"))}
            _ = v13
        }
    case -519:
        var v14 int64 = 9223372036854775807
        _ = v14
    case 2147483647:
        v2 += 2
    }
    if (!(true || false)) {
        var v15 int16 = 32767
        _ = v15
    }
    var v16 uint32 = 65535
    switch 32768 {
    case 4294967295:
        fmt.Println(v16)
    case 128:
        v16 = v16
    default:
        for fi17 := 0; fi17 < 3; fi17++ {
            fmt.Println(v16, v2)
            if (!(0.0 != float32((0.0 / -2.25)))) {
                v18 := GenS3{F0: "\n\t"}
                var v19 uint = 42
                v20 := GenS2{F0: genFn0("\n\t", 32767)}
                _ = v18
                _ = v19
                _ = v20
            }
            v16 += (v16 | v16)
            _ = fi17
        }
    }
    switch ("a" + ("\n\t" + "a")) {
    case "":
        _ = len(v3)
    case "a":
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 float64 = 1e10
            _ = fi21
            _ = v22
        }
    default:
        v23 := GenS1{F0: (v16 * (v16 >> ((255) & 63)))}
        _ = v23
    }
    if true {
        v24 := GenS1{F0: (v16 ^ v16)}
        switch -1 {
        case 9223372036854775807:
            var v25 string = "foo bar"
            _ = v25
        case -1:
            if (!(3.14159 >= 0.0)) {
                for fi26 := 0; fi26 < 15; fi26++ {
                    v2 = (v2 << ((404) & 63))
                    v16 -= (v16 ^ v16)
                    v2 -= 1
                    _ = fi26
                }
                if true {
                    var v27 int32 = -2147483648
                    var v28 float64 = 1.5
                    _ = v27
                    _ = v28
                }
                v16 = (v16 - 0)
            }
        default:
            fmt.Println(v2, v16)
        }
        _ = v24
    }
    if true {
        if false {
            switch "foo bar" {
            case "hello":
                var v29 bool = (1 <= -9223372036854775808)
                _ = v29
            case "foo bar":
                var v30 uint8 = 16
                _ = v30
            case "":
                v2 += 683
            }
            v16 = v16
        }
    } else {
        v2 = (v2 - v2)
        v31 := GenS4{F0: 2147483648}
        var v32 bool = (!(32767 <= 1))
        _ = v31
        _ = v32
    }
    if ("a" == ("hello" + (("" + "foo bar") + "foo bar"))) {
        if true {
            var v33 float64 = 1e10
            _ = v33
        }
    } else {
        switch ("\n\t" + "foo bar") {
        case "a":
            if (-128 > 127) {
                var v34 uint16 = 301
                v34 = v34
                _ = v34
            }
        default:
            for fi35 := 0; fi35 < 2; fi35++ {
                v36 := GenS5{F0: 0}
                _ = fi35
                _ = v36
            }
        }
        v1["hello"] = (!(!(0 <= 1)))
    }
    var v37 bool = false
    v38 := GenS6{F0: 335}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v38
}
