package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint16) uint16 {
    for fi5 := 0; fi5 < 17; fi5++ {
        p0 = (p0 >> ((1) & 63))
        var v6 uint64 = 0
        if (0 >= 9223372036854775807) {
            p0 += 1
            v7 := GenS1{F0: 42, F1: -128, F2: p0}
            v6 += 590
            _ = v7
        } else {
            switch "a" {
            case "":
                if (193 >= 2) {
                    var v8 string = (("foo bar" + ("hello" + (("hello" + "foo bar") + "\n\t"))) + "")
                    v9 := []int{(fi5 & fi5)}
                    _ = v8
                    _ = v9
                }
            case "hello":
                var v10 uint16 = uint16(v6)
                _ = v10
            case "\n\t":
                v6 = 0
            }
        }
        _ = fi5
        _ = v6
    }
    var v11 bool = false
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint64 = 1
        p0 -= (p0 << ((18446744073709551615) & 63))
        v14 := []int{fi12, int(p0)}
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 206, F1: 100, F2: 65535}
    v16 := []int{128}
    _ = p0
    _ = v11
    _ = v15
    _ = v16
    return p0
}
func genFn1(p0 int) uint64 {
    v17 := []int{(p0 << ((0) & 63)), 2147483647, -256, -2147483648}
    for fi18 := 0; fi18 < 16; fi18++ {
        p0 = fi18
        if (false || false) {
            p0 = (fi18 * fi18)
        }
        _ = fi18
    }
    if true {
        if (2147483647 > 2) {
            p0 = (p0 + (p0 / ((p0) | 1)))
            p0 += (p0 / ((p0) | 1))
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 string = (("a" + "hello") + "foo bar")
                _ = fi19
                _ = v20
            }
        } else {
            var v21 uint64 = 527
            v22 := GenS2{F0: 0, F1: 1}
            p0 = -306
            _ = v21
            _ = v22
        }
        p0 = 32768
        p0 = (p0 & (-(p0)))
    }
    var v23 uint64 = 42
    _ = p0
    _ = v17
    _ = v23
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2147483647
    var v3 uint = 255
    var v4 int64 = -1
    v4 = v4
    v24 := GenS3{F0: "\n\t"}
    v25 := v1[""]
    v3 -= (v3 - 1)
    fmt.Println(v3, v3)
    for fi26 := 0; fi26 < 9; fi26++ {
        fmt.Println(v25, v4, v25)
        for fi27 := 0; fi27 < 18; fi27++ {
            if (true || false) {
                if (255 <= 69) {
                    v2 += (^(v2))
                    if true {
                        fmt.Println(fi27, v3)
                        if ((v2 != v2) || true) {
                            delete(v1, "a")
                            var v28 int8 = 100
                            var v29 uint8 = 0
                            v25 += fi26
                            var v30 bool = true
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    }
                    if (255 <= 128) {
                        if ((100 != 25) && (((v2 >> ((v3) & 63)) >= 1) && (2147483647 >= 2147483647))) {
                            if (1 < -2147483648) {
                                var v31 uint32 = uint32(v3)
                                var v32 uint32 = v31
                                switch ("" + "a") {
                                case "hello":
                                    delete(v1, ("hello" + ""))
                                default:
                                    var v33 int32 = -2147483648
                                    _ = v33
                                }
                                _ = v31
                                _ = v32
                            }
                        } else {
                            v4 -= 42
                            if true {
                                v34 := GenS2{F0: 720, F1: 1}
                                fmt.Println(v25)
                                _ = v34
                            } else {
                                fmt.Println(v25, v3)
                                var v35 int8 = 42
                                v35 = v35
                                _ = v35
                            }
                            var v36 int16 = 42
                            _ = v36
                        }
                        v3 = (^(v3))
                    } else {
                        if false {
                            v37 := []int{}
                            var v38 uint = v3
                            v3 -= (^(v38))
                            _ = v37
                            _ = v38
                        } else {
                            v3 = 0
                        }
                        if true {
                            var v39 float32 = 3.14159
                            _ = v39
                        } else {
                            var v40 uint16 = 2
                            v1["\n\t"] = (v25 | int(v3))
                            _ = v40
                        }
                        v1[("a" + "a")] = (fi26 % ((v25) | 1))
                    }
                }
                v41 := v1["hello"]
                var v42 uint16 = genFn0(65535)
                _ = v41
                _ = v42
            } else {
                fmt.Println(v2)
                var v43 uint64 = 0
                v2 = (v2 ^ (^(v2)))
                _ = v43
            }
            var v44 uint16 = genFn0(32767)
            fmt.Println(v25)
            _ = fi27
            _ = v44
        }
        _ = fi26
    }
    v25 = v25
    v45 := GenS4{F0: 42}
    fmt.Println(v3, v4)
    for fi46 := 0; fi46 < 19; fi46++ {
        v47 := GenS1{F0: 66, F1: -66, F2: 770}
        switch 42 {
        case 2147483647:
            var v48 int16 = 1
            _ = v48
        default:
            fmt.Println(v25, v4)
        }
        _ = fi46
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v45
}
