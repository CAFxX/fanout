package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int8) uint16 {
    p0 = p0
    v5 := []int{}
    v6 := GenS1{F0: 65535}
    fmt.Println(p0)
    if len(v5) > 0 {
        v5[0] = 42
    }
    _ = p0
    _ = v5
    _ = v6
    return 628
}

func main() {
    var v1 uint = 4294967295
    v2 := []int{}
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    switch "foo bar" {
    case "foo bar":
        var v7 float64 = 0.0
        _ = v7
    case "\n\t":
        if (!(!(false || (694 > genFn0(-1))))) {
            v8 := v4[269]
            _ = v8
        } else {
            v1 = v1
        }
    default:
        fmt.Println(v1)
    }
    if (!((0.1 / float32(3.14159)) >= (-2.25 + 1.5))) {
        v9 := GenS1{F0: 1}
        _ = v9
    } else {
        if false {
            fmt.Println(v1)
            var v10 int32 = -2147483648
            if false {
                var v11 uint64 = 754
                _ = v11
            }
            _ = v10
        } else {
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 uint64 = 692
                var v14 string = (("hello" + "foo bar") + "hello")
                _ = fi12
                _ = v13
                _ = v14
            }
            switch -9223372036854775808 {
            case 127:
                v4[402] = 42
            }
        }
    }
    var v15 bool = (genFn0(-128) != 374)
    var v16 uint32 = 571
    _ = len(v2)
    fmt.Println(v1, v15, v1)
    fmt.Println(v16, v1)
    delete(v4, (v16 + (v16 - v16)))
    if (!false) {
        var v17 int8 = 107
        _ = v17
    } else {
        switch 32768 {
        case -2147483648:
            v18 := GenS1{F0: 953}
            _ = v18
        case 128:
            v15 = (202 >= 174)
        default:
            v19 := GenS1{F0: 550}
            _ = v19
        }
        if false {
            for fi20 := 0; fi20 < 13; fi20++ {
                for fi21 := 0; fi21 < 20; fi21++ {
                    if v15 {
                        fmt.Println(fi21)
                    }
                    v22 := GenS2{F0: (1000 <= 42)}
                    v23 := v4[v16]
                    if v15 {
                        var v24 float64 = 1e10
                        _ = v24
                    } else {
                        var v25 uint32 = v23
                        var v26 int32 = 14
                        if (2 <= 18446744073709551615) {
                            if (!false) {
                                switch fi21 {
                                case 9223372036854775807:
                                    if (!((v15 && true) && v15)) {
                                        fmt.Println(v15, fi21)
                                    }
                                case 128:
                                    fmt.Println(v26, fi20)
                                case 661:
                                    v16 += (^(v25))
                                default:
                                    switch fi20 {
                                    case 0:
                                        v27 := GenS2{F0: true}
                                        _ = v27
                                    case -1:
                                        v16 = v25
                                    }
                                }
                                v28 := GenS2{F0: (v1 < v1)}
                                _ = v28
                            } else {
                                fmt.Println(fi21, v26)
                            }
                            if true {
                                v15 = v15
                                if (0.0 <= 0.0) {
                                    var v29 int64 = 42
                                    _ = v29
                                }
                                if len(v2) > 0 {
                                    v2[0] = 128
                                }
                            } else {
                                if v15 {
                                    var v30 string = "\n\t"
                                    fmt.Println(fi21, fi21)
                                    _ = v30
                                } else {
                                    if false {
                                        v26 -= 248
                                        v31 := GenS1{F0: 2}
                                        _ = v31
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    var v32 bool = v15
                                    _ = v32
                                }
                                v33 := GenS1{F0: 255}
                                _ = v33
                            }
                        }
                        _ = v25
                        _ = v26
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v15 = (v1 <= (v1 - (v1 - 340)))
                var v34 int32 = 205
                _ = fi20
                _ = v34
            }
        }
    }
    v2 = append(v2, 127)
    _ = len(v4)
    if false {
        switch ("\n\t" + ("a" + ((("foo bar" + "hello") + "foo bar") + (("" + "hello") + (("\n\t" + "\n\t") + ("foo bar" + "hello")))))) {
        case "foo bar":
            v1 = (v1 & v1)
        default:
            fmt.Println(v16, v16)
        }
        fmt.Println(v15, v1)
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v15 = false
        fmt.Println(v1)
        _ = fi35
    }
    var v36 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
