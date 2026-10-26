package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int8; F2 int64 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 bool; F1 bool; F2 int32 }

func main() {
    v1 := []int{-530, -581, 23}
    var v2 uint16 = 32767
    v3 := []int{4294967295, 127, 128, 267}
    var v4 int64 = 1
    var v5 uint64 = 18446744073709551615
    v6 := make(map[string]int)
    switch 269 {
    case 799:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        var v7 int = 9223372036854775807
        _ = v7
    }
    if false {
        var v8 int16 = 32767
        v5 += 18446744073709551615
        var v9 uint32 = 65535
        _ = v8
        _ = v9
    } else {
        var v10 int16 = 120
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int = fi11
        var v13 int64 = v4
        _ = fi11
        _ = v12
        _ = v13
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    var v14 bool = ((!false) || true)
    for fi15 := 0; fi15 < 5; fi15++ {
        fmt.Println(v2, fi15, v2)
        var v16 uint = 65535
        fmt.Println(v16)
        _ = fi15
        _ = v16
    }
    if v14 {
        var v17 int32 = int32(v2)
        v17 += (v17 >> ((65535) & 63))
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := GenS1{F0: 0}
            var v20 uint8 = 128
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = v17
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 uint = 1
            var v24 bool = false
            _ = fi22
            _ = v23
            _ = v24
        }
        v5 = (v5 | (v5 / (((v5 % ((v5) | 1))) | 1)))
        _ = fi21
    }
    fmt.Println(v14)
    fmt.Println(v5, v14, v14)
    v2 = v2
    switch 1 {
    case 32768:
        v25 := GenS2{F0: 475, F1: 42, F2: (v4 | (v4 / ((v4) | 1)))}
        _ = v25
    case -2147483648:
        v14 = (1 >= 42)
    case -1:
        fmt.Println(v14)
    }
    v4 += (v4 & 1)
    switch 127 {
    case 128:
        for fi26 := 0; fi26 < 15; fi26++ {
            v2 -= v2
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 string = ("" + "\n\t")
                var v29 string = v28
                _ = fi27
                _ = v28
                _ = v29
            }
            for fi30 := 0; fi30 < 17; fi30++ {
                v14 = ((v5 / ((v5) | 1)) < v5)
                fmt.Println(fi30, fi30)
                if v14 {
                    _ = len(v1)
                    v5 = (v5 | v5)
                } else {
                    fmt.Println(fi26)
                }
                _ = len(v6)
                if (!v14) {
                    var v31 int16 = 392
                    v4 = v4
                    if len(v1) > 0 {
                        v1[0] = (fi26 >> ((65535) & 63))
                    }
                    _ = v31
                } else {
                    var v32 int = 0
                    if v14 {
                        switch (-(v32)) {
                        case -2147483648:
                            switch ("" + ((("foo bar" + "foo bar") + (("\n\t" + "foo bar") + "a")) + (("foo bar" + "\n\t") + ("hello" + ("\n\t" + "foo bar"))))) {
                            case "\n\t":
                                v33 := GenS2{F0: -97, F1: 127, F2: 9223372036854775807}
                                _ = v33
                            case "hello":
                                var v34 int64 = v4
                                var v35 int8 = -1
                                _ = v34
                                _ = v35
                            case "a":
                                var v36 int64 = -383
                                _ = v36
                            }
                        default:
                            v4 += (^(v4))
                        }
                        var v37 uint16 = v2
                        _ = v37
                    } else {
                        v38 := v6["\n\t"]
                        switch v32 {
                        case 42:
                            v3 = append(v3, (v32 / (((-(v38))) | 1)))
                        }
                        if ((v5 >> ((0) & 63)) > 0) {
                            var v39 uint32 = 2147483648
                            v14 = ((-11 == -128) || ((v5 * (v5 - v5)) == v5))
                            switch (v32 << ((4294967295) & 63)) {
                            case -2147483648:
                                var v40 uint64 = v5
                                _ = v40
                            case 4294967295:
                                var v41 float64 = 3.14159
                                _ = v41
                            case 2147483647:
                                v39 = (v39 - (v39 * (v39 | v39)))
                                var v42 float32 = 1e10
                                _ = v42
                            }
                            v43 := GenS3{F0: (v39 << ((2) & 63)), F1: fi30}
                            _ = v39
                            _ = v43
                        } else {
                            if len(v1) > 0 {
                                v1[0] = v38
                            }
                            var v44 uint = 65535
                            switch "" {
                            case "hello":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "":
                                if v14 {
                                    fmt.Println(v14, v4, v14)
                                    v32 += -2147483648
                                } else {
                                    switch (^(fi30)) {
                                    case 4294967295:
                                        var v45 int8 = 36
                                        _ = v45
                                    case -2147483648:
                                        delete(v6, "a")
                                        var v46 uint16 = (v2 & (v2 << (((^(v44))) & 63)))
                                        _ = v46
                                    case 965:
                                        _ = len(v1)
                                    }
                                    v32 = (-(fi26))
                                }
                            case "foo bar":
                                if (fi30 != fi30) {
                                    _ = len(v6)
                                    v6[(("\n\t" + "a") + "")] = (fi30 & fi30)
                                } else {
                                    fmt.Println(fi26)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = len(v3)
                                }
                                if (v4 >= v4) {
                                    var v47 int16 = 32767
                                    switch v38 {
                                    case 831:
                                        _ = len(v3)
                                    case -2147483648:
                                        var v48 float32 = ((3.14159 * float32(3.14159)) - 0.1)
                                        _ = v48
                                    case -174:
                                        var v49 bool = true
                                        _ = v49
                                    }
                                    _ = v47
                                } else {
                                    switch 127 {
                                    case 0:
                                        if v14 {
                                            switch fi26 {
                                            case 42:
                                                switch -2147483648 {
                                                case -2147483648:
                                                    v50 := GenS3{F0: uint32(v2), F1: v32}
                                                    _ = v50
                                                default:
                                                    var v51 int16 = 1
                                                    _ = v51
                                                }
                                            case 742:
                                                v4 += (v4 * -1)
                                            default:
                                                v44 = (v44 * 475)
                                            }
                                            v44 = (v44 + v44)
                                        }
                                    case 147:
                                        v52 := v6["hello"]
                                        _ = v52
                                    }
                                    v53 := GenS4{F0: false, F1: ("foo bar" >= "\n\t"), F2: 2147483647}
                                    v54 := v6[""]
                                    _ = v53
                                    _ = v54
                                }
                                var v55 bool = false
                                _ = v55
                            default:
                                v14 = v14
                            }
                            _ = v44
                        }
                        _ = v38
                    }
                    _ = v32
                }
                _ = fi30
            }
            _ = fi26
        }
    default:
        if (false && v14) {
            switch (("foo bar" + "a") + ("" + "foo bar")) {
            case "":
                var v56 int8 = -57
                _ = v56
            }
            var v57 int8 = 1
            _ = v57
        } else {
            var v58 string = "foo bar"
            _ = v58
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
