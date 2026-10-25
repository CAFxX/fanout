package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint64, p2 uint64) int16 {
    p2 = p1
    var v4 float32 = (0.1 + 1e10)
    p2 = p2
    var v5 int32 = -1
    var v6 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -2147483648
    var v3 int16 = 1
    switch ("foo bar" + ("hello" + ("hello" + ("\n\t" + "a")))) {
    case "foo bar":
        var v7 float64 = 1.5
        _ = v7
    case "":
        if false {
            fmt.Println(v2)
            switch 42 {
            case 4294967295:
                fmt.Println(v2, v3)
            }
            if (10 > 798) {
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 uint64 = 4294967295
                    switch (fi8 >> ((0) & 63)) {
                    case 1:
                        v10 := v1[2]
                        _ = v10
                    }
                    _ = fi8
                    _ = v9
                }
                switch "" {
                case "hello":
                    v11 := GenS1{F0: (false && (!(true || true))), F1: -867, F2: 127}
                    _ = v11
                case "":
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v2 = (v2 - 610)
                        _ = fi12
                    }
                }
            } else {
                for fi13 := 0; fi13 < 17; fi13++ {
                    var v14 int64 = 42
                    if (1 >= 2) {
                        if (726 == 1) {
                            v15 := GenS1{F0: (0 <= 18446744073709551615), F1: fi13, F2: 0}
                            _ = v15
                        } else {
                            v16 := GenS1{F0: ((v14 >> ((255) & 63)) < (v14 >> ((796) & 63))), F1: (fi13 & (fi13 >> ((42) & 63))), F2: fi13}
                            for fi17 := 0; fi17 < 6; fi17++ {
                                v18 := []int{fi17, 128, (^(fi17))}
                                var v19 int16 = (v3 + 32767)
                                if ((-(v3)) >= (v19 - v3)) {
                                    _ = len(v1)
                                } else {
                                    v20 := GenS1{F0: true, F1: fi17, F2: (fi17 >> ((1) & 63))}
                                    v14 = 1
                                    _ = v20
                                }
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            _ = v16
                        }
                        v21 := GenS1{F0: false, F1: (-(fi13)), F2: -784}
                        var v22 string = "foo bar"
                        _ = v21
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 8; fi23++ {
                            fmt.Println(v3, fi13, v3)
                            v24 := GenS1{F0: (((((float32((float64(3.14159) * ((1.5 - 3.14159) - float64((-2.25 / 1e10))))) / float32((0.1 + 0.1))) + 1.5) + 0.0) - -2.25) > -2.25), F1: fi23, F2: int(v14)}
                            if ((2 >= 42) && false) {
                                v14 += v14
                                v25 := GenS1{F0: false, F1: (fi23 | fi13), F2: 4294967295}
                                v26 := []int{(fi23 | fi23), fi23, (fi23 ^ (fi23 / (((fi23 << ((65535) & 63))) | 1))), fi23}
                                _ = v25
                                _ = v26
                            }
                            _ = fi23
                            _ = v24
                        }
                        if ((("a" + "a") + "a") != ("hello" + "foo bar")) {
                            v27 := GenS1{F0: (!(("\n\t" + "hello") != ("foo bar" + "a"))), F1: 9223372036854775807, F2: 32768}
                            for fi28 := 0; fi28 < 1; fi28++ {
                                var v29 uint64 = 2
                                _ = fi28
                                _ = v29
                            }
                            _ = v27
                        } else {
                            var v30 uint = 4294967295
                            switch ("a" + "hello") {
                            case "hello":
                                fmt.Println(v14)
                            case "foo bar":
                                var v31 uint = v30
                                _ = v31
                            default:
                                for fi32 := 0; fi32 < 18; fi32++ {
                                    v30 += (v30 << (((v30 ^ (v30 / (((v30 | v30)) | 1)))) & 63))
                                    _ = fi32
                                }
                            }
                            v3 = (v3 ^ (v3 * (v3 | v3)))
                            _ = v30
                        }
                        switch fi13 {
                        case 42:
                            if (-128 != -27) {
                                v3 = -1
                            } else {
                                fmt.Println(v3)
                            }
                        case 127:
                            fmt.Println(v2, v14)
                        }
                    }
                    v3 = genFn0((v14 - v14), 4294967295, 42)
                    _ = fi13
                    _ = v14
                }
                var v33 uint64 = 42
                v34 := GenS1{F0: true, F1: 9223372036854775807, F2: 128}
                _ = v33
                _ = v34
            }
        } else {
            v35 := GenS1{F0: true, F1: 0, F2: 42}
            _ = v35
        }
    case "\n\t":
        v2 = v2
    default:
        for fi36 := 0; fi36 < 16; fi36++ {
            fmt.Println(fi36)
            v37 := GenS1{F0: (true && (true || false)), F1: 127, F2: fi36}
            if (!(("hello" + "") == "foo bar")) {
                if false {
                    v38 := GenS2{F0: -9223372036854775808, F1: 42}
                    v2 -= (-(v2))
                    for fi39 := 0; fi39 < 13; fi39++ {
                        switch ("" + (((("hello" + "hello") + "\n\t") + "foo bar") + ("foo bar" + "foo bar"))) {
                        case "hello":
                            if false {
                                v1[734] = (-120 >= 0)
                                fmt.Println(fi39, v3, fi36)
                                if (0.1 < ((1.5 / 0.1) - 0.0)) {
                                    var v40 int8 = -1
                                    var v41 int64 = 1
                                    var v42 uint32 = 2
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                } else {
                                    v43 := GenS3{F0: 196}
                                    _ = v43
                                }
                            }
                        case "foo bar":
                            v44 := GenS1{F0: true, F1: (fi36 << ((0) & 63)), F2: (fi36 - fi36)}
                            _ = v44
                        case "\n\t":
                            switch (("foo bar" + "hello") + ((("\n\t" + (("foo bar" + (("" + "foo bar") + "foo bar")) + ("hello" + "hello"))) + "hello") + ("\n\t" + ("hello" + ("foo bar" + "foo bar"))))) {
                            case "hello":
                                v3 = (v3 & v3)
                            }
                        }
                        v2 += 894
                        v2 -= (^(v2))
                        _ = fi39
                    }
                    _ = v38
                } else {
                    var v45 int16 = v3
                    v46 := GenS2{F0: fi36, F1: 128}
                    _ = v45
                    _ = v46
                }
            } else {
                for fi47 := 0; fi47 < 2; fi47++ {
                    fmt.Println(v2)
                    _ = fi47
                }
                if ("\n\t" != "hello") {
                    v3 = -1
                } else {
                    var v48 uint64 = 2
                    v49 := GenS2{F0: (-(fi36)), F1: 255}
                    var v50 uint64 = v48
                    _ = v48
                    _ = v49
                    _ = v50
                }
                v51 := GenS3{F0: 0}
                _ = v51
            }
            _ = fi36
            _ = v37
        }
    }
    var v52 uint8 = 221
    switch "foo bar" {
    case "":
        v52 -= 127
    case "a":
        switch 0 {
        case 1:
            fmt.Println(v2, v52, v2)
        case -2147483648:
            v53 := GenS4{F0: 127}
            _ = v53
        }
    default:
        if true {
            v54 := []int{157, 32768}
            fmt.Println(v3, v3)
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 3; fi55++ {
        var v56 uint64 = 0
        v3 = 42
        v57 := v1[2147483647]
        _ = fi55
        _ = v56
        _ = v57
    }
    var v58 uint64 = 249
    var v59 int8 = int8(v52)
    fmt.Println(v58, v2)
    v3 = (v3 >> ((345) & 63))
    v1[806] = (false || true)
    v59 += -128
    var v60 int = 124
    var v61 uint = 255
    v62 := []int{(v60 & (v60 ^ 0)), -720, (v60 % (((v60 % ((v60) | 1))) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v58
    _ = v59
    _ = v60
    _ = v61
    _ = v62
}
