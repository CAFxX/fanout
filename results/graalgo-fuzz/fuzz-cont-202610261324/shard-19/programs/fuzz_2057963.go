package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 bool; F1 int64; F2 int64 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16; F1 uint8 }
func genFn0(p0 int64, p1 uint64) int8 {
    var v7 uint64 = (^(p1))
    p1 = 18446744073709551615
    var v8 int8 = 1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (v8 << ((65535) & 63))
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    var v3 int = 823
    v4 := make(map[int]int)
    var v5 int16 = -280
    v6 := []string{"\n\t"}
    switch v3 {
    case 42:
        var v9 uint32 = 0
        _ = v9
    case 0:
        var v10 uint64 = 4294967295
        _ = v10
    case -552:
        var v11 uint8 = 2
        _ = v11
    default:
        v12 := GenS1{F0: 0, F1: -1, F2: 2147483647}
        _ = v12
    }
    if v1 {
        v5 = v5
        v5 = 0
        if v1 {
            var v13 string = ("\n\t" + "\n\t")
            fmt.Println(v1, v5, v1)
            _ = v13
        }
    } else {
        var v14 bool = v1
        switch (("hello" + "") + "a") {
        case "a":
            var v15 int16 = -32768
            _ = v15
        case "foo bar":
            v5 = v5
        default:
            v3 -= (v3 << ((18446744073709551615) & 63))
        }
        _ = v14
    }
    var v16 int8 = genFn0(-9223372036854775808, 101)
    switch 42 {
    case 32768:
        if (!(127 != 242)) {
            if v1 {
                if (75 != 0) {
                    var v17 uint = 173
                    delete(v2, int(v16))
                    if true {
                        if false {
                            fmt.Println(v1, v1, v3)
                            var v18 int16 = (v5 / (((v5 | v5)) | 1))
                            if len(v6) > 0 {
                                v6[0] = (("\n\t" + ("a" + "foo bar")) + ("hello" + "\n\t"))
                            }
                            _ = v18
                        } else {
                            v17 -= 65535
                            v19 := GenS2{F0: 255, F1: 42}
                            v20 := GenS3{F0: 43}
                            _ = v19
                            _ = v20
                        }
                    }
                    _ = v17
                } else {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        var v22 uint8 = 42
                        if (!v1) {
                            var v23 uint16 = 65535
                            v24 := GenS4{F0: v1, F1: 18446744073709551615}
                            _ = v23
                            _ = v24
                        }
                        if (1.5 > 3.14159) {
                            v25 := GenS2{F0: 1, F1: (v22 + (v22 | v22))}
                            _ = v25
                        } else {
                            v26 := GenS4{F0: (!(1e10 <= (float32((1.5 / 3.14159)) / 1.0))), F1: 676}
                            v22 = 2
                            v5 = (v5 + (v5 + (v5 & (v5 / (((v5 / ((v5) | 1))) | 1)))))
                            _ = v26
                        }
                        _ = fi21
                        _ = v22
                    }
                }
                switch "hello" {
                case "foo bar":
                    if v1 {
                        for fi27 := 0; fi27 < 14; fi27++ {
                            var v28 string = "foo bar"
                            var v29 uint16 = 0
                            if (601 > (v29 + 1)) {
                                var v30 int64 = 1
                                v31 := GenS1{F0: 0, F1: (-(v16)), F2: 127}
                                _ = v30
                                _ = v31
                            } else {
                                _ = len(v6)
                                if (255 > 176) {
                                    for fi32 := 0; fi32 < 7; fi32++ {
                                        v5 -= v5
                                        _ = fi32
                                    }
                                }
                                fmt.Println(v1)
                            }
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    } else {
                        var v33 bool = v1
                        v33 = v1
                        _ = v33
                    }
                case "hello":
                    for fi34 := 0; fi34 < 8; fi34++ {
                        fmt.Println(v5)
                        fmt.Println(v3, v3, v16)
                        _ = fi34
                    }
                }
            }
            v35 := GenS5{F0: v1, F1: 0, F2: 819}
            _ = v35
        } else {
            v36 := GenS6{F0: "foo bar", F1: 42}
            if (float32(((0.0 * ((float64((1e10 + (1e10 - 3.14159))) / 3.14159) / (0.1 + 3.14159))) * ((1.5 * 0.1) + 0.0))) >= 1e10) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v3 += v3
                for fi37 := 0; fi37 < 14; fi37++ {
                    v16 += (v16 & 1)
                    v16 = -1
                    _ = fi37
                }
            }
            _ = v36
        }
    case 128:
        fmt.Println(v5)
    case 4294967295:
        for fi38 := 0; fi38 < 10; fi38++ {
            if (2147483647 >= -2147483648) {
                switch 128 {
                case 42:
                    var v39 float32 = 0.1
                    _ = v39
                case 0:
                    v40 := GenS2{F0: 0, F1: 139}
                    _ = v40
                case 9223372036854775807:
                    v5 -= v5
                }
                v3 -= -252
            }
            if v1 {
                if true {
                    v41 := GenS7{F0: 18446744073709551615, F1: 42}
                    _ = v41
                }
            }
            v42 := GenS8{F0: 0}
            _ = fi38
            _ = v42
        }
    }
    v16 -= (v16 << ((255) & 63))
    var v43 float64 = (3.14159 / 1.5)
    for fi44 := 0; fi44 < 1; fi44++ {
        v45 := GenS2{F0: 162, F1: 157}
        for fi46 := 0; fi46 < 14; fi46++ {
            v2[fi46] = fi46
            v3 = (fi46 << ((42) & 63))
            _ = fi46
        }
        if v1 {
            v47 := GenS9{F0: (-(v5)), F1: 127}
            if (!((!("" > "a")) && (2 != 65535))) {
                v16 = (^(v16))
            } else {
                v48 := GenS8{F0: 255}
                _ = v48
            }
            v3 -= fi44
            _ = v47
        }
        _ = fi44
        _ = v45
    }
    if (2147483648 <= 4294967295) {
        for fi49 := 0; fi49 < 12; fi49++ {
            v4[v3] = v3
            for fi50 := 0; fi50 < 12; fi50++ {
                delete(v4, (v3 - fi49))
                v16 = -113
                v16 = genFn0(0, 4294967295)
                _ = fi50
            }
            _ = fi49
        }
        for fi51 := 0; fi51 < 19; fi51++ {
            for fi52 := 0; fi52 < 19; fi52++ {
                var v53 float32 = (-2.25 / 0.1)
                _ = fi52
                _ = v53
            }
            _ = fi51
        }
    } else {
        fmt.Println(v43, v43, v1)
        v54 := v4[(v3 >> ((2) & 63))]
        for fi55 := 0; fi55 < 8; fi55++ {
            v5 -= v5
            _ = fi55
        }
        _ = v54
    }
    v3 += -687
    v3 += (v3 ^ (v3 - (v3 / ((v3) | 1))))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
