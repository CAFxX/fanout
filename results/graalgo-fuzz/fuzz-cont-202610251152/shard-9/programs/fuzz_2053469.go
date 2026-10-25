package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int64; F1 string; F2 uint8 }
type GenS4 struct { F0 int32; F1 int16; F2 int }
type GenS5 struct { F0 int; F1 int16; F2 int8 }
type GenS6 struct { F0 uint64; F1 int; F2 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 472
    var v3 float32 = (3.14159 * float32(float64((((0.1 / 0.1) - (0.1 * ((3.14159 * ((float32(1.5) - (float32(1e10) + ((3.14159 + 1.5) / 1.0))) * 3.14159)) - 0.0))) + float32(float64(1e10))))))
    var v4 uint32 = 42
    var v5 uint = 226
    v3 = (v3 - v3)
    var v6 int32 = -616
    if (669 <= -1) {
        var v7 string = "\n\t"
        v2 += v2
        v8 := GenS1{F0: v7}
        _ = v7
        _ = v8
    }
    v9 := []int{781, -878, -1}
    if (((!(0 > 0)) || true) && ((v6 % (((v6 - v6)) | 1)) >= v6)) {
        v10 := GenS2{F0: 42, F1: v4}
        v6 -= (v6 + v6)
        if len(v9) > 0 {
            v9[0] = 1
        }
        _ = v10
    } else {
        for fi11 := 0; fi11 < 14; fi11++ {
            if true {
                for fi12 := 0; fi12 < 11; fi12++ {
                    v4 = v4
                    _ = fi12
                }
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 int16 = int16(v4)
                    v14 = v14
                    _ = fi13
                    _ = v14
                }
            } else {
                fmt.Println(v3, v3, v2)
            }
            var v15 uint8 = 42
            switch (-(fi11)) {
            case 4294967295:
                switch (fi11 / (((^(fi11))) | 1)) {
                case 128:
                    if false {
                        v16 := GenS3{F0: 1, F1: "hello", F2: v15}
                        if (-1 == 813) {
                            var v17 float32 = v3
                            _ = v17
                        } else {
                            v5 = uint(v4)
                            for fi18 := 0; fi18 < 3; fi18++ {
                                fmt.Println(v4, fi18, v6)
                                _ = fi18
                            }
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        }
                        _ = v16
                    }
                case 1:
                    fmt.Println(v6, v2)
                case 0:
                    v6 += (^(v6))
                }
            default:
                v4 += 2
            }
            _ = fi11
            _ = v15
        }
    }
    v3 += v3
    var v19 uint64 = 4294967295
    v20 := GenS4{F0: (v6 + v6), F1: 1000, F2: 4294967295}
    v4 += (v4 - (v4 >> ((v5) & 63)))
    if (127 > -31) {
        v6 += (v6 - v6)
        if (v6 == int32(v4)) {
            var v21 int8 = 77
            v19 -= v19
            if len(v9) > 0 {
                v9[0] = 0
            }
            _ = v21
        } else {
            v2 += (v2 - v5)
            var v22 int16 = 1000
            for fi23 := 0; fi23 < 14; fi23++ {
                for fi24 := 0; fi24 < 8; fi24++ {
                    fmt.Println(v3, v2)
                    v25 := GenS1{F0: (("hello" + "\n\t") + "")}
                    v26 := GenS5{F0: fi23, F1: (v22 + 1000), F2: 100}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
            _ = v22
        }
        v5 = v2
    } else {
        fmt.Println(v2)
    }
    v27 := v1[185]
    switch "foo bar" {
    case "a":
        if true {
            switch ((("\n\t" + "a") + ("a" + "\n\t")) + "a") {
            case "\n\t":
                for fi28 := 0; fi28 < 11; fi28++ {
                    switch (v27 / ((fi28) | 1)) {
                    case 42:
                        v5 = (v2 >> ((v5) & 63))
                    case -712:
                        var v29 float32 = (v3 * 3.14159)
                        _ = v29
                    default:
                        v30 := GenS6{F0: v19, F1: (^(v27)), F2: 690}
                        _ = v30
                    }
                    var v31 int64 = 1
                    _ = len(v9)
                    _ = fi28
                    _ = v31
                }
            case "hello":
                for fi32 := 0; fi32 < 20; fi32++ {
                    var v33 uint32 = 65535
                    fmt.Println(v5)
                    _ = fi32
                    _ = v33
                }
            case "":
                v34 := GenS7{F0: v19}
                _ = v34
            }
            v1[128] = v27
            v35 := GenS8{F0: 100, F1: 32767, F2: 0}
            _ = v35
        } else {
            switch v27 {
            case 32768:
                var v36 float64 = 1e10
                _ = v36
            case -867:
                var v37 bool = false
                _ = v37
            case 1:
                if true {
                    _ = len(v1)
                } else {
                    if false {
                        fmt.Println(v4, v5)
                    }
                }
            }
            fmt.Println(v4, v4, v5)
            v6 -= (v6 << ((v2) & 63))
        }
    case "\n\t":
        switch ("hello" + ((("a" + "hello") + "hello") + "foo bar")) {
        case "":
            switch "hello" {
            case "":
                fmt.Println(v6, v3, v5)
            case "\n\t":
                var v38 int = int(v4)
                _ = v38
            default:
                switch "foo bar" {
                case "a":
                    var v39 int8 = 93
                    _ = v39
                }
            }
        case "hello":
            var v40 int = int(v5)
            _ = v40
        default:
            if true {
                if (!false) {
                    fmt.Println(v6, v3, v3)
                    v41 := GenS9{F0: v27}
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 2; fi42++ {
                        v9 = append(v9, 581)
                        _ = fi42
                    }
                    fmt.Println(v19, v4, v2)
                    switch 32768 {
                    case 273:
                        _ = len(v1)
                    case 32768:
                        var v43 float64 = (-2.25 + float64(v3))
                        _ = v43
                    default:
                        var v44 float32 = (1.5 * ((v3 * v3) * 3.14159))
                        _ = v44
                    }
                }
                for fi45 := 0; fi45 < 19; fi45++ {
                    switch "\n\t" {
                    case "a":
                        v6 = 1
                    case "foo bar":
                        v4 -= v4
                    }
                    v5 += v5
                    _ = fi45
                }
            }
        }
    case "hello":
        v9 = append(v9, (v27 * 128))
    default:
        var v46 float32 = 3.14159
        _ = v46
    }
    v47 := GenS9{F0: int(v6)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v9
}
