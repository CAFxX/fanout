package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int64; F1 uint; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 string; F1 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -2147483648
    v3 := make(map[uint32]uint32)
    v4 := []string{"hello", "\n\t", "\n\t"}
    var v5 int16 = -1
    var v6 string = "a"
    var v7 float32 = 1.5
    v8 := GenS1{F0: 0, F1: 4294967295}
    _ = len(v4)
    var v9 string = ("" + "")
    v5 += v5
    var v10 int = 32768
    fmt.Println(v7, v6)
    if (424 != 890) {
        if (float64(float32(1.5)) < 0.0) {
            if (-2.25 != 1e10) {
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 int16 = -32768
                    _ = fi11
                    _ = v12
                }
                var v13 int8 = 1
                _ = v13
            }
            v6 = ""
        }
        var v14 uint8 = 127
        v9 = ("a" + v9)
        _ = v14
    } else {
        v15 := GenS2{F0: 0, F1: 42, F2: 0}
        _ = v15
    }
    if (!false) {
        v16 := GenS1{F0: 127, F1: 65535}
        switch ((("\n\t" + "\n\t") + (("foo bar" + "") + v9)) + ("\n\t" + (v9 + v6))) {
        case "hello":
            if true {
                var v17 uint32 = 65535
                v9 = "\n\t"
                switch v10 {
                case 32768:
                    fmt.Println(v9, v17)
                default:
                    if true {
                        v4 = append(v4, v9)
                        delete(v1, 42)
                    } else {
                        for fi18 := 0; fi18 < 13; fi18++ {
                            if true {
                                v6 = v6
                                v19 := GenS1{F0: 100, F1: 42}
                                _ = v19
                            } else {
                                switch fi18 {
                                case 2147483647:
                                    v20 := GenS3{F0: v5}
                                    _ = v20
                                case 4294967295:
                                    var v21 bool = false
                                    _ = v21
                                default:
                                    _ = len(v4)
                                }
                            }
                            var v22 int8 = 1
                            var v23 uint64 = 428
                            _ = fi18
                            _ = v22
                            _ = v23
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                }
                _ = v17
            }
        default:
            v6 = v6
        }
        switch v10 {
        case -2147483648:
            var v24 bool = (92 < 127)
            _ = v24
        }
        _ = v16
    } else {
        switch int(v2) {
        case -9223372036854775808:
            switch (v6 + v9) {
            case "hello":
                if true {
                    fmt.Println(v5, v2, v7)
                    v5 += v5
                    v7 = (v7 + (float32((3.14159 * 0.0)) + v7))
                }
            case "":
                for fi25 := 0; fi25 < 20; fi25++ {
                    if (!(false || false)) {
                        fmt.Println(v2)
                        fmt.Println(v6, v9)
                    }
                    var v26 uint16 = 65535
                    _ = fi25
                    _ = v26
                }
            case "foo bar":
                fmt.Println(v5, v9, v10)
            }
        case -260:
            for fi27 := 0; fi27 < 14; fi27++ {
                fmt.Println(v6, v10)
                fmt.Println(v9, v2)
                v28 := GenS4{F0: 127, F1: v10}
                _ = fi27
                _ = v28
            }
        case -1:
            var v29 float32 = 1.5
            _ = v29
        }
        for fi30 := 0; fi30 < 6; fi30++ {
            fmt.Println(v10, fi30)
            var v31 uint64 = 2
            _ = fi30
            _ = v31
        }
    }
    if true {
        if true {
            var v32 bool = false
            fmt.Println(v5)
            _ = v32
        }
        v33 := GenS4{F0: -1, F1: v10}
        switch v6 {
        case "foo bar":
            var v34 int64 = 936
            _ = v34
        }
        _ = v33
    } else {
        v35 := GenS5{F0: "", F1: 2}
        v36 := GenS4{F0: 0, F1: v10}
        fmt.Println(v10, v6, v7)
        _ = v35
        _ = v36
    }
    v4 = append(v4, "a")
    v6 = (v6 + ("hello" + (("foo bar" + v6) + (("a" + ("foo bar" + v6)) + "a"))))
    v5 -= (v5 - v5)
    switch 486 {
    case 44:
        v37 := v3[65535]
        _ = v37
    case 32767:
        v38 := GenS3{F0: (v5 - int16(v2))}
        _ = v38
    default:
        if len(v4) > 0 {
            v4[0] = ""
        }
    }
    v4 = append(v4, ("\n\t" + (v9 + (v6 + "\n\t"))))
    for fi39 := 0; fi39 < 17; fi39++ {
        var v40 uint64 = 807
        _ = fi39
        _ = v40
    }
    v7 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
