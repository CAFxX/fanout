package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 int, p1 uint32) uint32 {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int = -912
        v9 := GenS1{F0: p1}
        v10 := []int{}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS1{F0: 2}
    p1 -= (p1 >> ((562) & 63))
    if true {
        p0 += 32767
        fmt.Println(p0)
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(p1)
            _ = fi12
        }
    } else {
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 uint32 = p1
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        v16 := []int{p0, (p0 + 42), (p0 + 1), (p0 >> ((65535) & 63))}
        _ = v16
    }
    var v17 int16 = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v17
    return p1
}
func genFn1(p0 uint, p1 int, p2 int64) uint64 {
    fmt.Println(p1, p2, p0)
    if false {
        v18 := GenS1{F0: 788}
        for fi19 := 0; fi19 < 2; fi19++ {
            p2 += 42
            switch ("foo bar" + "\n\t") {
            case "a":
                if (56 > -75) {
                    v20 := []int{fi19, fi19, (p1 << ((p0) & 63))}
                    var v21 int64 = p2
                    _ = v20
                    _ = v21
                }
            case "foo bar":
                var v22 uint8 = 42
                _ = v22
            default:
                var v23 uint = (p0 >> ((18446744073709551615) & 63))
                _ = v23
            }
            _ = fi19
        }
        _ = v18
    } else {
        var v24 uint = p0
        v25 := GenS1{F0: uint32(v24)}
        v24 = (p0 * p0)
        _ = v24
        _ = v25
    }
    p0 += 0
    fmt.Println(p1)
    switch ("hello" + ("hello" + ("" + ("hello" + ("\n\t" + "a"))))) {
    case "\n\t":
        p0 += (p0 * p0)
    case "foo bar":
        if true {
            p1 = -861
            if false {
                var v26 uint16 = 620
                _ = v26
            }
        } else {
            var v27 float32 = 0.1
            _ = v27
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 bool = (347 >= 127)
    v2 := []int{9223372036854775807, 2147483647, 438}
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]string)
    var v5 int32 = 1
    var v6 int32 = (v5 * v5)
    if v1 {
        var v28 float32 = 0.1
        var v29 int16 = 1
        _ = v28
        _ = v29
    } else {
        for fi30 := 0; fi30 < 15; fi30++ {
            for fi31 := 0; fi31 < 1; fi31++ {
                if true {
                    var v32 int = (fi31 >> ((94) & 63))
                    v33 := GenS2{F0: (v6 << ((210) & 63)), F1: 2147483647}
                    _ = v32
                    _ = v33
                }
                _ = fi31
            }
            v5 = v6
            _ = fi30
        }
    }
    var v34 uint32 = genFn0(0, 2147483648)
    v3 = v3
    for fi35 := 0; fi35 < 8; fi35++ {
        for fi36 := 0; fi36 < 6; fi36++ {
            v1 = v1
            switch fi35 {
            case 128:
                var v37 int = fi35
                _ = v37
            }
            if v1 {
                v34 = 4294967295
                v5 = v6
                v38 := GenS3{F0: "\n\t"}
                _ = v38
            }
            switch (fi36 & fi35) {
            case -2147483648:
                v5 += v5
            default:
                var v39 uint = 4294967295
                v3 = 472
                _ = v39
            }
            if v1 {
                var v40 int16 = int16(v34)
                _ = v40
            }
            _ = fi36
        }
        var v41 uint = 4294967295
        var v42 uint16 = 0
        _ = fi35
        _ = v41
        _ = v42
    }
    fmt.Println(v1, v6, v6)
    v6 += (v6 & (v6 + v6))
    var v43 int8 = 1
    v6 = v5
    switch "hello" {
    case "foo bar":
        v34 -= v34
    case "\n\t":
        if (!(244 >= 127)) {
            var v44 uint16 = 42
            v5 += (v6 & 412)
            _ = v44
        }
    case "":
        v3 = v3
    }
    switch "\n\t" {
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        var v45 int32 = v5
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
