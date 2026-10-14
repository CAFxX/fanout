package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint16; F2 string }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 uint64; F1 uint16; F2 uint32 }
func genFn0(p0 int32, p1 int32) int64 {
    v7 := GenS1{F0: 0}
    v8 := GenS1{F0: 0}
    var v9 float32 = -2.25
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -824
}

func main() {
    var v1 float64 = 0.0
    v2 := []int{-2147483648, 630, 32767, -211}
    var v3 uint64 = 2
    v4 := []int{}
    v5 := make(map[int]string)
    v6 := []uint32{407}
    for fi10 := 0; fi10 < 14; fi10++ {
        if false {
            if (!false) {
                switch ("a" + ("foo bar" + "foo bar")) {
                case "a":
                    v11 := GenS2{F0: 0}
                    _ = v11
                case "\n\t":
                    var v12 int32 = -2147483648
                    _ = v12
                case "hello":
                    for fi13 := 0; fi13 < 16; fi13++ {
                        if ("" < "") {
                            var v14 int = (fi10 | (fi10 & fi13))
                            v1 += v1
                            if true {
                                v3 = (v3 - 4294967295)
                                v15 := GenS2{F0: 42}
                                _ = v15
                            }
                            switch "foo bar" {
                            case "hello":
                                var v16 string = (("\n\t" + "a") + ("\n\t" + "foo bar"))
                                _ = v16
                            case "foo bar":
                                v17 := GenS1{F0: 127}
                                _ = v17
                            }
                            var v18 uint16 = 32767
                            _ = v14
                            _ = v18
                        }
                        v3 -= (v3 * (v3 >> ((255) & 63)))
                        fmt.Println(fi10, fi10)
                        var v19 float32 = float32(((float64(3.14159) + ((v1 + 0.0) + (v1 + float64(0.0)))) + float64(-2.25)))
                        _ = fi13
                        _ = v19
                    }
                default:
                    for fi20 := 0; fi20 < 15; fi20++ {
                        v21 := GenS3{F0: ((6 != 2) || false)}
                        _ = fi20
                        _ = v21
                    }
                }
                switch "" {
                case "a":
                    switch (fi10 << ((2) & 63)) {
                    case -308:
                        v1 -= v1
                    }
                case "\n\t":
                    var v22 bool = (false && (v1 != (v1 + v1)))
                    _ = v22
                case "":
                    if (v1 <= (v1 - v1)) {
                        v23 := v5[fi10]
                        _ = v23
                    }
                }
                switch fi10 {
                case 4294967295:
                    v3 += v3
                }
            }
            for fi24 := 0; fi24 < 19; fi24++ {
                switch fi24 {
                case -1:
                    var v25 float64 = v1
                    _ = v25
                }
                if (uint8(v3) < 1) {
                    v26 := v5[-103]
                    _ = v26
                }
                _ = fi24
            }
            var v27 uint64 = v3
            _ = v27
        } else {
            var v28 int = fi10
            v29 := v5[(v28 << ((49) & 63))]
            _ = v28
            _ = v29
        }
        _ = fi10
    }
    if ("\n\t" > "") {
        var v30 int32 = 727
        if (!true) {
            var v31 int64 = -1
            _ = v31
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = v30
    }
    v32 := GenS3{F0: false}
    for fi33 := 0; fi33 < 12; fi33++ {
        switch fi33 {
        case -2147483648:
            var v34 uint8 = 136
            _ = v34
        default:
            v35 := GenS4{F0: 246, F1: 286, F2: "hello"}
            _ = v35
        }
        v36 := GenS3{F0: (100 < 100)}
        _ = fi33
        _ = v36
    }
    v37 := GenS3{F0: (1e10 <= 1e10)}
    v1 += 1e10
    var v38 string = ""
    for fi39 := 0; fi39 < 15; fi39++ {
        fmt.Println(v38)
        v1 += 3.14159
        var v40 int = -535
        _ = fi39
        _ = v40
    }
    if true {
        _ = len(v5)
        var v41 float32 = (float32((float64(0.1) * float64(float32(0.1)))) - -2.25)
        switch -9223372036854775808 {
        case -627:
            v3 += (^(v3))
        case 0:
            _ = len(v5)
        case 32767:
            _ = len(v5)
        default:
            v3 -= v3
        }
        _ = v41
    } else {
        if (!(2147483647 < 42)) {
            switch "" {
            case "foo bar":
                switch (v38 + (v38 + ((v38 + (v38 + v38)) + v38))) {
                case "a":
                    var v42 int32 = 1
                    _ = v42
                case "foo bar":
                    v1 = v1
                case "":
                    v43 := GenS3{F0: (!((0 < 2) && (!true)))}
                    _ = v43
                default:
                    v44 := GenS4{F0: 432, F1: 32767, F2: v38}
                    _ = v44
                }
            case "\n\t":
                v45 := GenS5{F0: 867, F1: 1}
                _ = v45
            default:
                switch v38 {
                case "a":
                    _ = len(v4)
                case "\n\t":
                    var v46 string = (v38 + "hello")
                    _ = v46
                case "foo bar":
                    v47 := GenS4{F0: 540, F1: 2, F2: v38}
                    _ = v47
                default:
                    var v48 uint32 = 65535
                    _ = v48
                }
            }
            v49 := v5[128]
            _ = v49
        } else {
            v38 = v38
            v50 := GenS2{F0: 896}
            v51 := GenS6{F0: 2, F1: 538, F2: 710}
            _ = v50
            _ = v51
        }
    }
    v1 -= (v1 / ((v1 * -2.25) - 3.14159))
    if (9223372036854775807 > 1) {
        fmt.Println(v1, v3)
    } else {
        for fi52 := 0; fi52 < 1; fi52++ {
            fmt.Println(v38, v3, v38)
            v53 := GenS2{F0: 0}
            var v54 string = "a"
            _ = fi52
            _ = v53
            _ = v54
        }
        if (false && true) {
            fmt.Println(v38, v38)
            switch ((v38 + ("hello" + v38)) + v38) {
            case "\n\t":
                v55 := GenS1{F0: 96}
                _ = v55
            }
            fmt.Println(v1, v38, v1)
        } else {
            v1 += (v1 / v1)
            var v56 uint16 = 1
            var v57 uint16 = (v56 % ((v56) | 1))
            _ = v56
            _ = v57
        }
    }
    if (0 <= 349) {
        v38 = (v38 + "\n\t")
        var v58 int32 = -1
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
