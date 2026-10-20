package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 bool; F2 string }
type GenS6 struct { F0 uint32; F1 bool; F2 int64 }
func genFn0(p0 int16) int64 {
    var v4 int = -783
    var v5 string = "a"
    if (-2.25 == (-2.25 + (1.5 - 0.1))) {
        v6 := []int{v4, v4, v4}
        _ = v6
    }
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1() int {
    for fi7 := 0; fi7 < 1; fi7++ {
        switch fi7 {
        case 32767:
            var v8 int64 = genFn0(-1)
            _ = v8
        }
        if (((fi7 & (^(fi7))) < fi7) && true) {
            v9 := GenS1{F0: ""}
            _ = v9
        }
        v10 := GenS2{F0: 444}
        _ = fi7
        _ = v10
    }
    var v11 uint64 = 0
    v12 := GenS2{F0: 0}
    v11 += (v11 & (v11 % ((v11) | 1)))
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 string = "hello"
    var v2 uint32 = 2
    var v3 int8 = 1
    var v13 float32 = float32((1e10 - 3.14159))
    switch v1 {
    case "hello":
        var v14 int64 = 693
        _ = v14
    case "a":
        for fi15 := 0; fi15 < 12; fi15++ {
            v3 = (v3 - (v3 - 96))
            v16 := []int{genFn1(), fi15}
            _ = fi15
            _ = v16
        }
    case "":
        var v17 uint8 = 2
        _ = v17
    default:
        switch v1 {
        case "foo bar":
            v1 = "foo bar"
        case "a":
            if (false && ((v2 >> ((18446744073709551615) & 63)) < v2)) {
                var v18 uint32 = v2
                v3 -= 65
                v19 := GenS1{F0: "hello"}
                _ = v18
                _ = v19
            }
        default:
            var v20 int = 32767
            _ = v20
        }
    }
    switch v1 {
    case "\n\t":
        v13 = ((((float32(1e10) / ((float32(0.1) * v13) * ((v13 / (0.1 / v13)) * (v13 / v13)))) + float32(0.1)) + 1e10) - (v13 + 3.14159))
    case "hello":
        v1 = (v1 + v1)
    }
    switch v1 {
    case "":
        var v21 uint = 0
        _ = v21
    }
    var v22 int32 = 0
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            v22 += v22
            var v25 int64 = 42
            if true {
                v26 := GenS2{F0: v25}
                v27 := GenS1{F0: ""}
                _ = v26
                _ = v27
            }
            v28 := GenS3{F0: v2}
            fmt.Println(fi24)
            fmt.Println(v13)
            v29 := GenS4{F0: 1}
            _ = fi24
            _ = v25
            _ = v28
            _ = v29
        }
        fmt.Println(fi23, v22, v2)
        _ = fi23
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        v2 = v2
        _ = fi30
    }
    switch v1 {
    case "a":
        v3 += -1
    default:
        switch "" {
        case "":
            fmt.Println(v2)
        case "\n\t":
            switch "\n\t" {
            case "foo bar":
                v2 = v2
            case "\n\t":
                fmt.Println(v22)
            }
        case "hello":
            v22 += 2147483647
        }
    }
    v31 := []int{32768, 953, -504}
    var v32 int8 = (v3 >> ((848) & 63))
    if false {
        var v33 string = (v1 + v1)
        v1 = "\n\t"
        _ = v33
    } else {
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 uint = 722
            fmt.Println(v35, v32, v1)
            v13 = 0.0
            _ = fi34
            _ = v35
        }
        v1 = "foo bar"
    }
    var v36 int16 = -32768
    var v37 float32 = v13
    var v38 uint32 = v2
    switch -1 {
    case 0:
        var v39 uint64 = 0
        _ = v39
    case -102:
        for fi40 := 0; fi40 < 3; fi40++ {
            for fi41 := 0; fi41 < 3; fi41++ {
                if true {
                    switch (fi41 & fi41) {
                    case 127:
                        v42 := GenS5{F0: 750, F1: (true && true), F2: (("hello" + (("hello" + v1) + "\n\t")) + v1)}
                        _ = v42
                    }
                    if (v37 <= (v37 + 1.5)) {
                        v43 := GenS6{F0: v38, F1: ((v13 + v37) < ((0.0 + v37) - v37)), F2: 0}
                        var v44 uint16 = 32767
                        var v45 string = v1
                        _ = v43
                        _ = v44
                        _ = v45
                    } else {
                        var v46 uint = 1
                        v47 := GenS2{F0: 0}
                        _ = v46
                        _ = v47
                    }
                } else {
                    var v48 uint64 = 725
                    var v49 uint = 4294967295
                    v50 := GenS4{F0: 2}
                    _ = v48
                    _ = v49
                    _ = v50
                }
                _ = fi41
            }
            fmt.Println(v3)
            v22 -= v22
            _ = fi40
        }
    case 32767:
        v51 := GenS5{F0: 255, F1: (32768 < 32767), F2: v1}
        _ = v51
    default:
        var v52 int = 1
        _ = v52
    }
    v13 = (v37 * (1.5 * v37))
    v37 += 3.14159
    if false {
        var v53 float32 = float32(-2.25)
        _ = v53
    } else {
        var v54 bool = (42 > (v36 - -808))
        v54 = (!(((genFn1() > 42) && (244 >= genFn1())) && (!(!true))))
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v37
    _ = v38
}
