package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int64; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 bool; F2 bool }
func genFn0(p0 uint32) string {
    v4 := GenS1{F0: 416}
    if false {
        for fi5 := 0; fi5 < 3; fi5++ {
            v6 := []int{}
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 int64 = 42
                p0 += (p0 + 2)
                fmt.Println(fi5, p0)
                _ = fi7
                _ = v8
            }
            _ = fi5
            _ = v6
        }
        p0 = (^(p0))
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 bool = (("a" + "") <= "\n\t")
        v11 := GenS1{F0: 2}
        fmt.Println(fi9, p0, p0)
        _ = fi9
        _ = v10
        _ = v11
    }
    p0 = (p0 << ((1) & 63))
    _ = p0
    _ = v4
    return "a"
}
func genFn1() int64 {
    switch -1 {
    case 128:
        var v12 uint = 691
        _ = v12
    default:
        var v13 int32 = -2147483648
        _ = v13
    }
    var v14 uint16 = 32767
    fmt.Println(v14)
    var v15 int = -2147483648
    v16 := []int{(v15 ^ 127), v15, (^(v15)), v15}
    _ = v14
    _ = v15
    _ = v16
    return 9223372036854775807
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]string)
    v3 := make(map[string]bool)
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int8 = 73
        var v19 string = "foo bar"
        var v20 int16 = v1
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        v1 += v1
        v1 = (-(v1))
        _ = fi21
    }
    v22 := GenS2{F0: -1, F1: -340, F2: 0}
    fmt.Println(v1)
    for fi23 := 0; fi23 < 11; fi23++ {
        switch fi23 {
        case -1:
            for fi24 := 0; fi24 < 8; fi24++ {
                v25 := GenS3{F0: 1000}
                v26 := GenS3{F0: v1}
                _ = fi24
                _ = v25
                _ = v26
            }
        case 2147483647:
            v1 = (v1 + v1)
        case -2147483648:
            if (!true) {
                v27 := GenS2{F0: 331, F1: 358, F2: (fi23 >> ((2) & 63))}
                _ = v27
            } else {
                v1 -= (v1 / ((32767) | 1))
                for fi28 := 0; fi28 < 20; fi28++ {
                    fmt.Println(fi23)
                    var v29 int32 = 1
                    _ = fi28
                    _ = v29
                }
            }
        }
        _ = fi23
    }
    if (!(255 <= 42)) {
        var v30 int16 = v1
        switch "hello" {
        case "hello":
            for fi31 := 0; fi31 < 17; fi31++ {
                if true {
                    switch "" {
                    case "hello":
                        var v32 int = fi31
                        _ = v32
                    default:
                        v30 = v1
                    }
                    v1 += (v1 + (v30 - v1))
                    var v33 string = "hello"
                    _ = v33
                }
                var v34 int32 = 42
                var v35 uint = 700
                _ = fi31
                _ = v34
                _ = v35
            }
        case "foo bar":
            v36 := GenS3{F0: 1}
            _ = v36
        case "\n\t":
            switch (((("" + "a") + "") + "foo bar") + "") {
            case "foo bar":
                fmt.Println(v1, v30)
            case "a":
                v37 := GenS4{F0: 1}
                _ = v37
            case "\n\t":
                v1 = v30
            }
        }
        var v38 uint = 2
        _ = v30
        _ = v38
    } else {
        v1 -= v1
        for fi39 := 0; fi39 < 18; fi39++ {
            if ((1e10 < 3.14159) && ((v1 - v1) >= v1)) {
                delete(v2, 42)
                v1 += (^(v1))
                var v40 uint8 = 1
                _ = v40
            }
            _ = fi39
        }
    }
    var v41 uint32 = 0
    switch 77 {
    case 1:
        v41 -= v41
    case 128:
        switch (("hello" + ("" + "a")) + "foo bar") {
        case "":
            fmt.Println(v1, v1)
        }
    case 450:
        var v42 int8 = -1
        _ = v42
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        v44 := GenS3{F0: v1}
        _ = fi43
        _ = v44
    }
    v45 := GenS1{F0: 0}
    switch "foo bar" {
    case "\n\t":
        switch ("\n\t" + "foo bar") {
        case "hello":
            v1 = (-(v1))
        }
    case "":
        switch "\n\t" {
        case "a":
            var v46 int16 = v1
            _ = v46
        case "\n\t":
            switch 2147483647 {
            case 127:
                switch int(v41) {
                case 9223372036854775807:
                    v41 += v41
                case 32768:
                    v47 := GenS5{F0: true, F1: (true && true), F2: true}
                    _ = v47
                default:
                    var v48 int16 = v1
                    _ = v48
                }
            case 1:
                var v49 uint32 = 784
                _ = v49
            }
        case "foo bar":
            fmt.Println(v41)
        default:
            fmt.Println(v41)
        }
    default:
        var v50 int16 = v1
        _ = v50
    }
    v1 -= v1
    v41 = v41
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v41
    _ = v45
}
