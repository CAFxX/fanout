package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int64, p1 bool) uint {
    if (!true) {
        switch 964 {
        case 127:
            var v4 uint16 = 0
            _ = v4
        case 128:
            p1 = (!false)
        case 598:
            v5 := []int{-9223372036854775808, 2147483647, -737}
            _ = v5
        default:
            if false {
                if p1 {
                    var v6 uint64 = 2
                    var v7 uint32 = 21
                    _ = v6
                    _ = v7
                }
                v8 := GenS1{F0: 70}
                if (0 != 18446744073709551615) {
                    fmt.Println(p1, p1)
                } else {
                    v9 := GenS1{F0: 381}
                    _ = v9
                }
                _ = v8
            } else {
                var v10 float32 = 1.5
                v11 := GenS1{F0: 182}
                _ = v10
                _ = v11
            }
        }
    } else {
        switch (("hello" + "\n\t") + ("a" + "a")) {
        case "\n\t":
            v12 := []int{-9223372036854775808, 32767, 0, 32767}
            _ = v12
        }
        switch (("foo bar" + ("hello" + "")) + ("\n\t" + "foo bar")) {
        case "hello":
            var v13 uint64 = 4294967295
            _ = v13
        default:
            if ((0.1 - 1.5) > (((0.1 * (3.14159 / (float32(0.1) - 3.14159))) / 1.0) / 1.0)) {
                v14 := []int{}
                p1 = (100 == -25)
                _ = v14
            }
        }
        p1 = p1
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := []int{(fi15 % (((fi15 * (fi15 * 128))) | 1)), fi15, 128}
        p1 = (!true)
        if true {
            var v17 uint32 = 135
            p1 = (int32(v17) >= 863)
            _ = v17
        } else {
            var v18 uint8 = 140
            var v19 int8 = 127
            var v20 int8 = v19
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    var v21 int32 = 1
    var v22 string = ""
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return 65535
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -1
    v3 := make(map[uint32]string)
    v23 := GenS1{F0: 32767}
    var v24 int32 = 1
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 int32 = (v24 | (v24 & v24))
        fmt.Println(v24)
        for fi27 := 0; fi27 < 3; fi27++ {
            fmt.Println(v26, v2, fi25)
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    v2 = -1
    var v28 string = "hello"
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 string = ("" + ((v28 + (v28 + ((("\n\t" + "") + "hello") + (("\n\t" + (v28 + (v28 + ((((((v28 + v28) + ((v28 + v28) + "hello")) + "\n\t") + (v28 + v28)) + "") + v28)))) + v28)))) + "\n\t"))
        switch fi29 {
        case 686:
            var v31 float64 = -2.25
            _ = v31
        case 1:
            var v32 int8 = -1
            _ = v32
        }
        v30 = v28
        _ = fi29
        _ = v30
    }
    if true {
        v33 := GenS2{F0: 1, F1: v2, F2: v2}
        if (true || false) {
            if true {
                v34 := []int{-2147483648, 80, -610}
                _ = v34
            } else {
                fmt.Println(v28)
                fmt.Println(v28, v24)
                v35 := []int{}
                _ = v35
            }
            var v36 bool = false
            v1[9223372036854775807] = 53
            _ = v36
        }
        _ = v33
    } else {
        v37 := GenS3{F0: 0}
        var v38 int32 = v24
        v38 += 0
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 uint64 = 2
        for fi41 := 0; fi41 < 8; fi41++ {
            switch (fi39 - 2147483647) {
            case 42:
                var v42 float64 = float64(float32(float64(1e10)))
                _ = v42
            case -523:
                v40 += v40
            case -9223372036854775808:
                v28 = v28
            default:
                if (!(!false)) {
                    var v43 uint32 = 4294967295
                    var v44 int8 = 1
                    var v45 uint8 = 0
                    _ = v43
                    _ = v44
                    _ = v45
                }
            }
            v40 = (v40 >> ((65535) & 63))
            _ = fi41
        }
        _ = fi39
        _ = v40
    }
    if (326 < 2) {
        _ = len(v3)
    }
    var v46 int32 = v24
    fmt.Println(v2, v2, v28)
    if (-2147483648 > 732) {
        var v47 uint64 = 18446744073709551615
        _ = v47
    }
    var v48 int8 = 127
    var v49 int32 = (v46 - (-(v24)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v46
    _ = v48
    _ = v49
}
