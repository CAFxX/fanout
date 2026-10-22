package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int64; F2 uint64 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int32 = -1
    var v3 uint8 = 1
    var v4 uint16 = 255
    var v5 uint16 = 277
    var v6 bool = (2147483648 < 337)
    v4 -= 255
    v6 = (-560 == 0)
    v7 := GenS1{F0: (v6 || ((v1 << ((271) & 63)) <= -9223372036854775808)), F1: 1, F2: v6}
    switch "" {
    case "hello":
        if v6 {
            v1 = (^(v1))
            if v6 {
                switch 1 {
                case 32767:
                    fmt.Println(v3, v6, v6)
                case -265:
                    v1 = 0
                case -1:
                    v8 := []int{1}
                    _ = v8
                default:
                    v3 += (v3 ^ (v3 * 42))
                }
                for fi9 := 0; fi9 < 15; fi9++ {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        if (4294967295 > 65535) {
                            if v6 {
                                var v11 string = "\n\t"
                                _ = v11
                            } else {
                                var v12 bool = ((-(v1)) > v1)
                                _ = v12
                            }
                            var v13 float32 = 0.1
                            _ = v13
                        }
                        fmt.Println(v2, v3, v1)
                        _ = fi10
                    }
                    v1 += int64(v5)
                    v14 := []int{fi9, 32768}
                    _ = fi9
                    _ = v14
                }
            }
            v5 = uint16(v3)
        } else {
            v3 -= v3
            for fi15 := 0; fi15 < 8; fi15++ {
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 int32 = v2
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
            for fi18 := 0; fi18 < 19; fi18++ {
                v19 := GenS2{F0: 0}
                _ = fi18
                _ = v19
            }
        }
    case "\n\t":
        for fi20 := 0; fi20 < 2; fi20++ {
            switch fi20 {
            case -1:
                v21 := GenS1{F0: (-128 >= 42), F1: 128, F2: v6}
                _ = v21
            case 128:
                switch ("" + ("foo bar" + "\n\t")) {
                case "a":
                    v22 := GenS1{F0: true, F1: v3, F2: false}
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 10; fi23++ {
                        switch (fi23 & fi20) {
                        case 128:
                            v24 := []int{fi23}
                            _ = v24
                        case 32767:
                            if v6 {
                                v1 += 435
                            }
                        }
                        switch ("\n\t" + ("" + (("foo bar" + "hello") + ("\n\t" + ("a" + ""))))) {
                        case "foo bar":
                            switch (fi23 & 4294967295) {
                            case 1:
                                if v6 {
                                    v6 = (-244 != -699)
                                    v6 = v6
                                } else {
                                    var v25 bool = false
                                    v26 := GenS3{F0: 127, F1: 0, F2: 1}
                                    if false {
                                        var v27 string = "a"
                                        var v28 int = (fi20 / ((fi23) | 1))
                                        v29 := []int{}
                                        fmt.Println(v4, v25)
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    } else {
                                        switch int(v5) {
                                        case -274:
                                            v30 := []int{fi20, (fi20 + fi20), (fi20 & (-(fi23))), (fi23 ^ fi23)}
                                            _ = v30
                                        }
                                        v31 := []int{}
                                        _ = v31
                                    }
                                    _ = v25
                                    _ = v26
                                }
                            }
                        case "\n\t":
                            v32 := GenS1{F0: v6, F1: (v3 | v3), F2: (!v6)}
                            _ = v32
                        }
                        _ = fi23
                    }
                }
            case 32768:
                if (!true) {
                    v33 := GenS1{F0: (100 != 100), F1: v3, F2: false}
                    _ = v33
                }
            }
            for fi34 := 0; fi34 < 4; fi34++ {
                fmt.Println(v6, v3, v2)
                fmt.Println(v4, v1, v4)
                v4 = (v4 / ((375) | 1))
                v5 -= 65535
                _ = fi34
            }
            var v35 uint32 = 65535
            _ = fi20
            _ = v35
        }
    default:
        v36 := GenS1{F0: v6, F1: (v3 * (v3 * v3)), F2: (0 >= v1)}
        _ = v36
    }
    v37 := []int{1, -462, 32768}
    v38 := GenS2{F0: 1000}
    fmt.Println(v5, v3)
    v3 += (v3 | v3)
    if (false && ((("foo bar" > "a") && true) && (!true))) {
        switch (("foo bar" + ("hello" + "hello")) + ("hello" + ("a" + "\n\t"))) {
        case "":
            fmt.Println(v2, v3)
        case "hello":
            var v39 int64 = v1
            _ = v39
        case "foo bar":
            v5 += v4
        default:
            v2 += v2
        }
    }
    v40 := GenS1{F0: v6, F1: (v3 - v3), F2: v6}
    v37 = append(v37, 32768)
    for fi41 := 0; fi41 < 2; fi41++ {
        var v42 uint64 = 803
        var v43 int8 = 100
        var v44 int8 = v43
        _ = fi41
        _ = v42
        _ = v43
        _ = v44
    }
    v2 = (^(v2))
    var v45 bool = (-1 > 128)
    var v46 bool = (1.5 > ((1.5 / 1.0) / -2.25))
    for fi47 := 0; fi47 < 4; fi47++ {
        fmt.Println(v3, v6)
        _ = fi47
    }
    var v48 int64 = v1
    fmt.Println(v6, v46, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v45
    _ = v46
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
