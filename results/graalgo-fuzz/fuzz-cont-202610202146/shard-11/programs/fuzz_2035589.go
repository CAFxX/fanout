package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint64; F2 uint }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64 }
func genFn0() int32 {
    var v5 uint64 = 1
    v6 := GenS1{F0: 718, F1: 42}
    v5 -= v5
    var v7 uint64 = v5
    _ = v5
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 string) int8 {
    p0 = ""
    p0 = "\n\t"
    _ = p0
    return 1
}

func main() {
    v1 := []uint32{2, 215, 65535}
    v2 := []string{("foo bar" + (("\n\t" + "foo bar") + ("hello" + "")))}
    v3 := make(map[uint32]string)
    var v4 int16 = -32768
    _ = len(v2)
    if (true || (!(2147483647 >= 4294967295))) {
        if (false && (false || (0.1 < 1.5))) {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
        } else {
            v8 := GenS1{F0: -116, F1: 990}
            _ = v8
        }
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 float64 = 0.0
            delete(v3, 460)
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 7; fi11++ {
            switch fi11 {
            case -569:
                if len(v2) > 0 {
                    v2[0] = ""
                }
            default:
                v4 = (v4 - v4)
            }
            switch ("a" + "\n\t") {
            case "hello":
                fmt.Println(v4)
            case "foo bar":
                if true {
                    v4 = 1000
                    v4 = v4
                } else {
                    v4 = (v4 / (((v4 - 804)) | 1))
                    v12 := GenS1{F0: 0, F1: (fi11 & fi11)}
                    if (false || true) {
                        for fi13 := 0; fi13 < 12; fi13++ {
                            fmt.Println(v4, fi13)
                            if true {
                                v4 += (v4 | (v4 + 42))
                            }
                            _ = fi13
                        }
                        var v14 int8 = 67
                        _ = v14
                    } else {
                        switch "hello" {
                        case "":
                            for fi15 := 0; fi15 < 5; fi15++ {
                                v4 += (v4 / (((v4 << ((255) & 63))) | 1))
                                _ = fi15
                            }
                        case "foo bar":
                            v4 = 1000
                        default:
                            if false {
                                delete(v3, 2147483648)
                            }
                        }
                        v16 := GenS1{F0: fi11, F1: (fi11 | int(v4))}
                        _ = v16
                    }
                    _ = v12
                }
            }
            v4 = v4
            _ = fi11
        }
    } else {
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 uint64 = 18446744073709551615
            if true {
                var v19 float32 = (0.1 + (3.14159 * 1.5))
                v20 := GenS2{F0: 255}
                var v21 int16 = v4
                _ = v19
                _ = v20
                _ = v21
            } else {
                v18 -= v18
                for fi22 := 0; fi22 < 16; fi22++ {
                    v18 += v18
                    v3[2147483648] = "hello"
                    _ = fi22
                }
                if ((0 <= 1000) && (!(-9223372036854775808 > 0))) {
                    v23 := GenS3{F0: (("foo bar" + ("\n\t" + "\n\t")) + "hello"), F1: 42, F2: 556}
                    if (v4 != v4) {
                        v18 = 2
                    }
                    if (!true) {
                        switch 32767 {
                        case -177:
                            for fi24 := 0; fi24 < 12; fi24++ {
                                fmt.Println(fi24, v18)
                                var v25 uint8 = 42
                                _ = fi24
                                _ = v25
                            }
                        case -9223372036854775808:
                            fmt.Println(fi17)
                        case 1:
                            if false {
                                var v26 int32 = 397
                                _ = v26
                            } else {
                                var v27 bool = false
                                v4 = (v4 | v4)
                                _ = v27
                            }
                        }
                    } else {
                        for fi28 := 0; fi28 < 9; fi28++ {
                            var v29 uint8 = 127
                            _ = fi28
                            _ = v29
                        }
                        var v30 uint64 = v18
                        _ = v30
                    }
                    _ = v23
                }
            }
            _ = fi17
            _ = v18
        }
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 float32 = (-2.25 + 0.0)
        _ = fi31
        _ = v32
    }
    if true {
        var v33 bool = true
        v4 = 946
        for fi34 := 0; fi34 < 3; fi34++ {
            v33 = v33
            _ = fi34
        }
        _ = v33
    } else {
        v4 += (v4 >> ((65535) & 63))
    }
    v35 := GenS4{F0: "hello", F1: 2}
    if (!(-9223372036854775808 <= -1)) {
        fmt.Println(v4)
        _ = len(v3)
    } else {
        if false {
            v4 = (v4 & v4)
            for fi36 := 0; fi36 < 4; fi36++ {
                v37 := GenS5{F0: (("foo bar" + "\n\t") + (("foo bar" + "foo bar") + "foo bar"))}
                switch ("foo bar" + "hello") {
                case "foo bar":
                    for fi38 := 0; fi38 < 19; fi38++ {
                        var v39 float64 = 0.0
                        v40 := GenS1{F0: fi36, F1: fi36}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                case "a":
                    if false {
                        v4 = v4
                        v41 := GenS6{F0: (true || (!(1 >= 255)))}
                        _ = v41
                    }
                case "hello":
                    fmt.Println(v4, v4)
                }
                fmt.Println(v4, fi36)
                _ = fi36
                _ = v37
            }
            for fi42 := 0; fi42 < 2; fi42++ {
                v4 = v4
                _ = fi42
            }
        } else {
            switch 32768 {
            case -236:
                var v43 int64 = -1
                _ = v43
            case 9223372036854775807:
                v44 := GenS5{F0: "a"}
                _ = v44
            case 4294967295:
                v4 = (^(v4))
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v4 = v4
            for fi45 := 0; fi45 < 16; fi45++ {
                var v46 string = ""
                v4 = v4
                for fi47 := 0; fi47 < 3; fi47++ {
                    v46 = "foo bar"
                    v46 = ("\n\t" + v46)
                    _ = fi47
                }
                _ = fi45
                _ = v46
            }
        }
        v4 += (v4 / ((-1) | 1))
    }
    v48 := GenS7{F0: 103}
    var v49 uint8 = 42
    v50 := GenS5{F0: "hello"}
    switch (("hello" + "foo bar") + "a") {
    case "foo bar":
        v4 = (v4 * v4)
    }
    v49 -= v49
    v49 = (^(v49))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v48
    _ = v49
    _ = v50
}
