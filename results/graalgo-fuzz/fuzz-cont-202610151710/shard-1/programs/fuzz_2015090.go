package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 int16 }
func genFn0(p0 uint16) int {
    if true {
        p0 = 42
        var v5 uint8 = 127
        _ = v5
    } else {
        switch "" {
        case "foo bar":
            p0 -= p0
        case "a":
            var v6 string = "a"
            _ = v6
        case "\n\t":
            var v7 uint8 = 42
            _ = v7
        }
        switch "\n\t" {
        case "":
            if false {
                for fi8 := 0; fi8 < 8; fi8++ {
                    var v9 uint8 = 2
                    v10 := GenS1{F0: v9}
                    var v11 float32 = 3.14159
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                var v12 bool = true
                var v13 uint8 = 1
                _ = v12
                _ = v13
            }
        case "\n\t":
            for fi14 := 0; fi14 < 2; fi14++ {
                for fi15 := 0; fi15 < 14; fi15++ {
                    v16 := []int{(fi15 >> ((18446744073709551615) & 63)), 42}
                    var v17 bool = (65535 > 4294967295)
                    var v18 int64 = 530
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 bool = true
                    v21 := GenS2{F0: p0, F1: 42, F2: p0}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = fi14
            }
        default:
            var v22 uint8 = 0
            _ = v22
        }
    }
    if false {
        for fi23 := 0; fi23 < 2; fi23++ {
            v24 := GenS2{F0: p0, F1: 42, F2: p0}
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            v26 := GenS2{F0: p0, F1: 289, F2: p0}
            _ = fi25
            _ = v26
        }
    } else {
        for fi27 := 0; fi27 < 2; fi27++ {
            v28 := []int{fi27, (fi27 ^ 9223372036854775807), (fi27 ^ (fi27 ^ fi27))}
            if (true || false) {
                for fi29 := 0; fi29 < 8; fi29++ {
                    v30 := GenS3{F0: 0, F1: 42}
                    switch (-(fi27)) {
                    case 128:
                        fmt.Println(fi27, fi29)
                    }
                    _ = fi29
                    _ = v30
                }
            } else {
                var v31 int = int(p0)
                _ = len(v28)
                fmt.Println(fi27, fi27, fi27)
                _ = v31
            }
            _ = fi27
            _ = v28
        }
    }
    _ = p0
    return 1
}
func genFn1(p0 int8, p1 string, p2 bool) uint {
    fmt.Println(p1, p0)
    p1 = "foo bar"
    switch 2147483647 {
    case 874:
        if false {
            var v32 uint8 = 255
            _ = v32
        } else {
            v33 := GenS4{F0: 42}
            _ = v33
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 uint64 = 883
    var v2 int16 = 42
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    v2 += v2
    v2 = (v2 ^ 864)
    if false {
        v2 = v2
    } else {
        for fi34 := 0; fi34 < 19; fi34++ {
            fmt.Println(v2)
            _ = fi34
        }
        if true {
            for fi35 := 0; fi35 < 18; fi35++ {
                v1 -= (v1 << ((4294967295) & 63))
                _ = fi35
            }
        }
        var v36 string = (("" + ("" + "foo bar")) + "\n\t")
        _ = v36
    }
    fmt.Println(v2)
    v1 -= v1
    var v37 int16 = v2
    if false {
        v2 = 87
    } else {
        v4[""] = (0 == 316)
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        for fi39 := 0; fi39 < 11; fi39++ {
            if true {
                v1 += 18446744073709551615
            }
            v2 -= (v37 << ((255) & 63))
            _ = fi39
        }
        switch "a" {
        case "":
            var v40 uint8 = 5
            _ = v40
        case "a":
            switch "\n\t" {
            case "":
                v41 := []int{}
                _ = v41
            case "hello":
                var v42 uint32 = 2147483647
                _ = v42
            }
        default:
            fmt.Println(fi38)
        }
        _ = fi38
    }
    var v43 int = 128
    var v44 uint16 = 255
    v2 = v37
    switch ("foo bar" + "\n\t") {
    case "hello":
        var v45 uint64 = v1
        _ = v45
    case "foo bar":
        var v46 int8 = -128
        _ = v46
    case "":
        v2 -= (v2 ^ 0)
    default:
        v47 := GenS5{F0: "hello", F1: 408}
        _ = v47
    }
    var v48 float32 = ((1e10 - 1e10) / (float32(1e10) * 3.14159))
    if true {
        v37 = (v2 | v2)
    } else {
        var v49 uint32 = 1
        v1 = 812
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v43
    _ = v44
    _ = v48
}
