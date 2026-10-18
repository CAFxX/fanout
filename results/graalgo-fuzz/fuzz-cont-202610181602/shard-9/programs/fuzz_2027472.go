package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 bool; F1 bool; F2 int64 }
func genFn0() int32 {
    var v7 uint32 = 2147483647
    v7 -= v7
    _ = v7
    return -435
}
func genFn1(p0 int64) uint64 {
    for fi8 := 0; fi8 < 10; fi8++ {
        p0 += p0
        v9 := GenS1{F0: (65535 >= 2), F1: 2147483647}
        var v10 int64 = (p0 - (p0 + (p0 - 42)))
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int8 = 13
    _ = p0
    _ = v11
    return 1
}

func main() {
    var v1 float32 = 0.0
    v2 := []uint32{68}
    var v3 string = ("a" + "")
    v4 := []string{"", "a", v3}
    var v5 int16 = 568
    var v6 int64 = -719
    switch 127 {
    case -908:
        for fi12 := 0; fi12 < 11; fi12++ {
            v1 = -2.25
            var v13 uint32 = 2
            _ = fi12
            _ = v13
        }
    case 42:
        if (1 == -1) {
            v3 = v3
            v3 = v3
            fmt.Println(v3)
        }
    default:
        v5 = (v5 % (((-(v5))) | 1))
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        fmt.Println(v3, v6)
        _ = fi14
    }
    var v15 int = 1
    switch 42 {
    case 117:
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 uint8 = 255
            _ = fi16
            _ = v17
        }
    case -155:
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 int8 = 13
            v3 = "hello"
            _ = fi18
            _ = v19
        }
    case 4294967295:
        var v20 uint64 = 1
        _ = v20
    default:
        if false {
            if false {
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 uint32 = 65535
                    _ = fi21
                    _ = v22
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = (v3 + v3)
                }
                var v23 int32 = -2147483648
                for fi24 := 0; fi24 < 20; fi24++ {
                    for fi25 := 0; fi25 < 2; fi25++ {
                        switch ((v3 + (v3 + "a")) + v3) {
                        case "":
                            fmt.Println(fi24, v1)
                        case "a":
                            v26 := GenS2{F0: (true || true), F1: true, F2: v6}
                            switch (v3 + v3) {
                            case "hello":
                                v27 := GenS2{F0: (!(39 == 1)), F1: ((true || false) || (v23 < 1)), F2: int64(fi25)}
                                _ = v27
                            case "\n\t":
                                if false {
                                    var v28 int64 = (v6 & v6)
                                    _ = v28
                                }
                            case "a":
                                v6 -= -9223372036854775808
                            }
                            if (false || true) {
                                var v29 string = ((((v3 + v3) + "\n\t") + v3) + ("foo bar" + ""))
                                if len(v4) > 0 {
                                    v4[0] = v3
                                }
                                v15 = fi24
                                _ = v29
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = "a"
                                }
                                var v30 uint16 = 80
                                _ = v30
                            }
                            var v31 bool = true
                            _ = v26
                            _ = v31
                        case "\n\t":
                            v5 = v5
                        default:
                            v3 = v3
                            v5 = v5
                        }
                        var v32 float32 = v1
                        if false {
                            v4 = append(v4, v3)
                            fmt.Println(v23)
                            v15 -= fi25
                        }
                        _ = fi25
                        _ = v32
                    }
                    _ = fi24
                }
                _ = v23
            }
        } else {
            fmt.Println(v6)
            fmt.Println(v3, v3)
        }
    }
    v5 -= v5
    var v33 bool = ((v3 + ("foo bar" + "")) <= ((v3 + "") + (v3 + v3)))
    for fi34 := 0; fi34 < 2; fi34++ {
        v35 := GenS2{F0: v33, F1: (v33 && (v33 && (v6 == v6))), F2: v6}
        var v36 bool = ((v33 && v33) || (1 > 0))
        _ = fi34
        _ = v35
        _ = v36
    }
    if len(v4) > 0 {
        v4[0] = v3
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 bool = v33
        if (3.14159 != (float32((1e10 - 1e10)) + v1)) {
            var v39 uint64 = 42
            var v40 int = fi37
            _ = len(v2)
            _ = v39
            _ = v40
        } else {
            _ = len(v4)
            var v41 uint = 0
            v4 = append(v4, (v3 + "\n\t"))
            _ = v41
        }
        v2 = append(v2, 2)
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
