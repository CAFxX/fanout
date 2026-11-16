package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int64 = -1
    var v2 int32 = -2147483648
    var v3 uint = 65535
    v4 := []uint32{0, 65535, 65535}
    v5 := make(map[string]uint32)
    if true {
        v1 += 0
        v3 += (v3 * 4294967295)
    } else {
        v2 -= v2
        fmt.Println(v3, v2, v1)
    }
    if ((0 < 44) && false) {
        var v6 bool = true
        var v7 float32 = -2.25
        v8 := GenS1{F0: 127}
        _ = v6
        _ = v7
        _ = v8
    } else {
        switch 42 {
        case 127:
            switch 1 {
            case 0:
                v3 = v3
            case 4294967295:
                for fi9 := 0; fi9 < 20; fi9++ {
                    var v10 int64 = 42
                    _ = fi9
                    _ = v10
                }
            default:
                v1 = v1
            }
        }
    }
    _ = len(v4)
    if (-9223372036854775808 == 4294967295) {
        v11 := v5["\n\t"]
        if ("" != (("\n\t" + ("hello" + "")) + "\n\t")) {
            var v12 uint = v3
            for fi13 := 0; fi13 < 3; fi13++ {
                if (!false) {
                    v14 := GenS1{F0: 102}
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        fmt.Println(fi15, v2)
                        var v16 string = ("a" + ("\n\t" + ""))
                        _ = fi15
                        _ = v16
                    }
                    v17 := GenS1{F0: 0}
                    _ = v17
                }
                _ = fi13
            }
            v18 := v5["hello"]
            _ = v12
            _ = v18
        } else {
            var v19 uint64 = 0
            for fi20 := 0; fi20 < 20; fi20++ {
                v3 -= v3
                _ = fi20
            }
            for fi21 := 0; fi21 < 20; fi21++ {
                v4 = append(v4, (^(v11)))
                for fi22 := 0; fi22 < 6; fi22++ {
                    v19 -= v19
                    if ((int64(v11) == (v1 >> ((v3) & 63))) || (0.1 >= -2.25)) {
                        if len(v4) > 0 {
                            v4[0] = v11
                        }
                        var v23 int = (fi22 * (fi21 << ((v3) & 63)))
                        var v24 uint64 = 2
                        _ = v23
                        _ = v24
                    }
                    if true {
                        fmt.Println(v2, fi21, v11)
                        if true {
                            if true {
                                var v25 int16 = 1
                                _ = v25
                            } else {
                                var v26 uint = (v3 - 4294967295)
                                v3 = (v26 - 18446744073709551615)
                                var v27 float64 = 0.0
                                v3 = v3
                                fmt.Println(fi22, v3, v27)
                                var v28 float64 = v27
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                            switch (fi21 - fi21) {
                            case 187:
                                if true {
                                    v19 -= (v19 & (v19 % ((v19) | 1)))
                                }
                                v19 += v19
                            case 158:
                                var v29 uint = v3
                                _ = v29
                            case 41:
                                var v30 int = fi21
                                _ = v30
                            default:
                                v31 := v5["hello"]
                                _ = v31
                            }
                        }
                        if (!false) {
                            if true {
                                fmt.Println(v11, v11)
                                if (false || true) {
                                    delete(v5, "")
                                } else {
                                    if false {
                                        v1 += v1
                                    }
                                    v32 := GenS1{F0: uint8(v19)}
                                    _ = v32
                                }
                                switch "a" {
                                case "":
                                    fmt.Println(v1, v11, fi21)
                                case "foo bar":
                                    var v33 float64 = 0.1
                                    _ = v33
                                case "\n\t":
                                    var v34 bool = ((0.1 > (0.1 * (3.14159 * 1e10))) && true)
                                    _ = v34
                                default:
                                    v11 = v11
                                }
                            } else {
                                if (!true) {
                                    if false {
                                        switch "foo bar" {
                                        case "hello":
                                            var v35 string = "hello"
                                            _ = v35
                                        case "foo bar":
                                            switch 32768 {
                                            case 4294967295:
                                                switch ("hello" + "a") {
                                                case "hello":
                                                    v36 := GenS2{F0: 0}
                                                    if false {
                                                        if len(v4) > 0 {
                                                            v4[0] = v11
                                                        }
                                                        v1 = 42
                                                        v3 -= (^(v3))
                                                    } else {
                                                        var v37 int32 = v2
                                                        var v38 string = "hello"
                                                        v19 += uint64(v3)
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    _ = v36
                                                case "foo bar":
                                                    v39 := GenS2{F0: 255}
                                                    _ = v39
                                                default:
                                                    var v40 float64 = 0.0
                                                    _ = v40
                                                }
                                            case 2147483647:
                                                v41 := v5["hello"]
                                                _ = v41
                                            case -365:
                                                if len(v4) > 0 {
                                                    v4[0] = 1
                                                }
                                            default:
                                                v42 := GenS2{F0: 65535}
                                                _ = v42
                                            }
                                        }
                                        var v43 uint16 = 558
                                        fmt.Println(v1)
                                        v3 = v3
                                        fmt.Println(v11, v19, fi21)
                                        _ = v43
                                    }
                                    var v44 uint16 = 42
                                    if (!true) {
                                        v19 = v19
                                        var v45 int = fi21
                                        var v46 int16 = 0
                                        var v47 float32 = (float32((3.14159 / -2.25)) + 3.14159)
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    }
                                    _ = v44
                                }
                                var v48 uint16 = 255
                                _ = v48
                            }
                            v49 := v5["hello"]
                            _ = v49
                        } else {
                            if (false && false) {
                                v19 = v19
                                v5["a"] = v11
                                v11 += 2147483648
                            } else {
                                var v50 float64 = 0.1
                                _ = v50
                            }
                            var v51 uint32 = (v11 ^ 156)
                            _ = v51
                        }
                    }
                    _ = fi22
                }
                _ = fi21
            }
            _ = v19
        }
        _ = v11
    } else {
        _ = len(v5)
        if true {
            switch "a" {
            case "":
                fmt.Println(v3, v3)
            case "foo bar":
                if true {
                    var v52 string = ("a" + "a")
                    v53 := v5["foo bar"]
                    _ = v52
                    _ = v53
                }
            }
        }
        var v54 string = ("a" + "a")
        _ = v54
    }
    fmt.Println(v3, v2)
    v2 -= v2
    v5[(("" + (("" + "") + "foo bar")) + ("hello" + "hello"))] = 65535
    for fi55 := 0; fi55 < 7; fi55++ {
        var v56 int16 = int16(v1)
        v1 = v1
        fmt.Println(fi55)
        _ = fi55
        _ = v56
    }
    var v57 bool = (false || (100 == -122))
    var v58 int64 = (v1 >> (((v3 / (((v3 + 18446744073709551615)) | 1))) & 63))
    if (!v57) {
        v3 = 42
        fmt.Println(v58)
        var v59 bool = true
        _ = v59
    }
    v3 -= (v3 | (v3 & (v3 >> (((^(v3))) & 63))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
}
