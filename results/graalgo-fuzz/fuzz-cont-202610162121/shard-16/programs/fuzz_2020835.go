package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 string; F1 int32; F2 string }
type GenS3 struct { F0 int16 }
func genFn0(p0 int16) string {
    p0 = 32767
    var v4 uint8 = 1
    if true {
        switch "" {
        case "foo bar":
            fmt.Println(p0, p0)
        default:
            switch "a" {
            case "hello":
                var v5 int64 = 249
                _ = v5
            case "":
                for fi6 := 0; fi6 < 11; fi6++ {
                    v4 -= (v4 << ((255) & 63))
                    _ = fi6
                }
            case "\n\t":
                v7 := []int{52, 831}
                _ = v7
            }
        }
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    switch 9223372036854775807 {
    case 980:
        v4 = (v4 << ((885) & 63))
    case -1:
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int16 = 32
            _ = fi10
            _ = v11
        }
    case 1:
        var v12 uint32 = 2147483648
        _ = v12
    default:
        var v13 int8 = 0
        _ = v13
    }
    _ = p0
    _ = v4
    return "a"
}
func genFn1(p0 uint32) int32 {
    var v14 uint64 = 42
    var v15 uint8 = 2
    switch "" {
    case "foo bar":
        switch "a" {
        case "":
            p0 = (p0 | (p0 * 2147483648))
        case "hello":
            var v16 uint32 = (p0 | p0)
            _ = v16
        case "foo bar":
            v17 := GenS1{F0: 42, F1: 255}
            _ = v17
        }
    case "":
        v14 -= 42
    case "a":
        var v18 int16 = 0
        _ = v18
    default:
        var v19 bool = true
        _ = v19
    }
    var v20 bool = false
    v15 -= 255
    _ = p0
    _ = v14
    _ = v15
    _ = v20
    return 0
}

func main() {
    var v1 int8 = 1
    v2 := []string{"a", "\n\t"}
    var v3 string = "a"
    var v21 int32 = 42
    var v22 float32 = 3.14159
    if (int(v1) < 9223372036854775807) {
        fmt.Println(v1, v22)
    } else {
        v22 = v22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        switch ((v3 + v3) + v3) {
        case "foo bar":
            var v24 uint16 = 65535
            _ = v24
        default:
            switch (("\n\t" + v3) + "\n\t") {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "a":
                var v25 uint16 = 2
                _ = v25
            case "\n\t":
                for fi26 := 0; fi26 < 9; fi26++ {
                    fmt.Println(v22, fi26, v3)
                    _ = fi26
                }
            }
        }
        _ = fi23
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        fmt.Println(v3)
        if true {
            v28 := GenS2{F0: v3, F1: (v21 * genFn1(1)), F2: "\n\t"}
            _ = v28
        }
        for fi29 := 0; fi29 < 18; fi29++ {
            fmt.Println(v21, v1)
            var v30 uint8 = 42
            v22 = float32((float64(v22) * float64(v22)))
            _ = fi29
            _ = v30
        }
        _ = fi27
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        switch v3 {
        case "":
            for fi32 := 0; fi32 < 3; fi32++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi32
            }
        case "a":
            var v33 uint8 = 1
            _ = v33
        }
        _ = len(v2)
        var v34 uint8 = 127
        _ = fi31
        _ = v34
    }
    if (4294967295 == 86) {
        var v35 int8 = 15
        v36 := GenS3{F0: 231}
        v3 = (v3 + (("" + v3) + (v3 + "a")))
        _ = v35
        _ = v36
    }
    var v37 string = "hello"
    var v38 bool = true
    if v38 {
        if v38 {
            v22 = 0.0
            v37 = "a"
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v38
}
