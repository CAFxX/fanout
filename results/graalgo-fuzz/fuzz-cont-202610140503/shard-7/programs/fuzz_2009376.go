package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }
type GenS4 struct { F0 int8; F1 int8; F2 int }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 1.5
    var v3 int = 128
    var v4 uint16 = 255
    var v5 uint = 531
    delete(v1, "\n\t")
    switch "\n\t" {
    case "\n\t":
        v4 = 736
    case "hello":
        v6 := []int{128, v3, v3}
        _ = v6
    }
    if true {
        var v7 int16 = -32768
        _ = v7
    } else {
        var v8 uint8 = 170
        if ((2 == 444) && false) {
            v1["a"] = (v3 | -2147483648)
            if (-1 < 2147483647) {
                v3 = (v3 << (((v5 - 65535)) & 63))
                var v9 int32 = 0
                v3 -= 418
                _ = v9
            } else {
                var v10 uint64 = 4294967295
                v11 := GenS1{F0: (0 > 2147483647)}
                _ = v10
                _ = v11
            }
        }
        _ = v8
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS2{F0: -48}
        for fi14 := 0; fi14 < 2; fi14++ {
            if false {
                switch ((("hello" + "") + "") + "foo bar") {
                case "":
                    v15 := []int{289, (fi12 % (((-(fi14))) | 1)), fi14, -387}
                    _ = v15
                case "foo bar":
                    var v16 uint64 = 18446744073709551615
                    _ = v16
                case "\n\t":
                    v17 := GenS1{F0: (uint64(v5) > 18446744073709551615)}
                    _ = v17
                default:
                    var v18 uint = (^(v5))
                    _ = v18
                }
                v19 := GenS2{F0: 1}
                _ = v19
            }
            v5 += v5
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    switch (v3 >> (((v5 + (v5 << ((289) & 63)))) & 63)) {
    case 42:
        var v20 int32 = 2147483647
        _ = v20
    case 1:
        v2 = (3.14159 * float64(1e10))
    }
    switch v3 {
    case 9223372036854775807:
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 uint = v5
            for fi23 := 0; fi23 < 16; fi23++ {
                v3 -= fi23
                v4 -= (v4 >> ((4294967295) & 63))
                fmt.Println(v22, fi23)
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    }
    switch v3 {
    case -2147483648:
        if (!(!((6 != 2147483648) && true))) {
            var v24 float64 = 3.14159
            _ = v24
        }
    case 42:
        v5 += (v5 << ((v5) & 63))
    default:
        fmt.Println(v2)
    }
    v25 := []int{(v3 & (-(v3))), v3, v3, v3}
    v26 := GenS3{F0: true, F1: (-(v3)), F2: 4294967295}
    var v27 uint8 = 39
    switch "a" {
    case "":
        for fi28 := 0; fi28 < 20; fi28++ {
            v25 = append(v25, -520)
            for fi29 := 0; fi29 < 20; fi29++ {
                v27 += (^(v27))
                v30 := GenS4{F0: 1, F1: 100, F2: (v3 >> ((137) & 63))}
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    case "\n\t":
        if len(v25) > 0 {
            v25[0] = -9223372036854775808
        }
    case "foo bar":
        var v31 uint = uint(v4)
        _ = v31
    }
    fmt.Println(v4)
    if (true || (32767 == 32767)) {
        switch "a" {
        case "foo bar":
            v5 -= 0
        case "hello":
            fmt.Println(v2)
        default:
            v3 -= (v3 & (v3 << (((v5 ^ 255)) & 63)))
        }
        var v32 int32 = 42
        _ = v32
    }
    if false {
        var v33 uint8 = 2
        v27 = 0
        _ = v33
    } else {
        for fi34 := 0; fi34 < 8; fi34++ {
            v5 = v5
            _ = fi34
        }
        for fi35 := 0; fi35 < 2; fi35++ {
            var v36 int32 = 231
            v37 := GenS2{F0: -1}
            v36 -= v36
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    v38 := v1["a"]
    switch v38 {
    case 42:
        v4 = 65535
    case 2147483647:
        switch (^(v38)) {
        case 350:
            v2 += v2
        default:
            if true {
                _ = len(v25)
            }
        }
    case 731:
        for fi39 := 0; fi39 < 4; fi39++ {
            var v40 string = ("hello" + "foo bar")
            v4 = (^(v4))
            switch fi39 {
            case 47:
                fmt.Println(fi39, v2)
            case 377:
                var v41 int8 = 100
                _ = v41
            }
            _ = fi39
            _ = v40
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
