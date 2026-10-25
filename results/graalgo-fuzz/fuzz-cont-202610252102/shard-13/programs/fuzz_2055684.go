package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int8 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    v7 := GenS1{F0: (((((("hello" + (("" + ("\n\t" + "")) + ("\n\t" + "\n\t"))) + "") + ("foo bar" + "foo bar")) + (((("foo bar" + (("a" + (("foo bar" + ("foo bar" + "a")) + ("foo bar" + "hello"))) + "\n\t")) + "a") + "hello") + ("foo bar" + ("" + ("foo bar" + (("a" + "a") + ("a" + ("" + ("\n\t" + "foo bar"))))))))) + "a") + "\n\t")}
    var v8 uint64 = 0
    if (true || (!true)) {
        v9 := GenS1{F0: ""}
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := []int{(fi10 / ((2147483647) | 1)), fi10}
            var v12 int64 = -609
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            v8 += (v8 * v8)
            _ = fi13
        }
        _ = v9
    }
    v8 = 18446744073709551615
    _ = v7
    _ = v8
    return 461
}
func genFn1() bool {
    for fi14 := 0; fi14 < 15; fi14++ {
        if (true || false) {
            var v15 uint = 0
            if false {
                var v16 bool = ((fi14 << ((255) & 63)) >= fi14)
                v17 := []int{fi14}
                _ = v16
                _ = v17
            }
            _ = v15
        } else {
            var v18 uint8 = 0
            _ = v18
        }
        _ = fi14
    }
    var v19 uint16 = 1
    _ = v19
    return false
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{42, 2147483647, 2}
    var v4 uint = 4294967295
    var v5 int8 = -1
    var v20 uint8 = 193
    v21 := GenS2{F0: true}
    var v22 float64 = (3.14159 - (0.0 * 0.1))
    delete(v1, -2147483648)
    var v23 uint8 = uint8(v2)
    if (42 > -1) {
        for fi24 := 0; fi24 < 11; fi24++ {
            fmt.Println(v23, v5)
            v5 += (v5 & (v5 ^ (v5 - v5)))
            _ = fi24
        }
    } else {
        if ("a" <= (("a" + ("\n\t" + (("" + ((("hello" + (("a" + "") + ("hello" + "a"))) + ("a" + ("" + ("a" + (("foo bar" + ("\n\t" + "hello")) + "\n\t"))))) + "hello")) + "a"))) + "")) {
            switch ((("\n\t" + "foo bar") + "a") + "foo bar") {
            case "a":
                v25 := GenS1{F0: ""}
                _ = v25
            case "":
                var v26 uint16 = uint16(v4)
                _ = v26
            case "foo bar":
                var v27 int32 = 320
                _ = v27
            default:
                v20 = 244
            }
            if true {
                v28 := GenS3{F0: (^(v2)), F1: (^(v20)), F2: 0}
                delete(v1, 1)
                var v29 int64 = -1
                _ = v28
                _ = v29
            } else {
                _ = len(v1)
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v30 float64 = v22
        fmt.Println(v22, v5, v30)
        _ = v30
    }
    if (v5 != 42) {
        switch (("a" + "\n\t") + "hello") {
        case "foo bar":
            for fi31 := 0; fi31 < 11; fi31++ {
                v4 += (v4 % (((v4 / (((v4 & 42)) | 1))) | 1))
                var v32 float32 = 0.0
                v33 := GenS4{F0: v5}
                _ = fi31
                _ = v32
                _ = v33
            }
        case "":
            for fi34 := 0; fi34 < 3; fi34++ {
                var v35 uint64 = (^(v2))
                v4 = 436
                _ = fi34
                _ = v35
            }
        case "\n\t":
            fmt.Println(v23, v20)
        }
        v4 += v4
    }
    fmt.Println(v5, v2, v20)
    var v36 int16 = genFn0()
    switch 127 {
    case -2147483648:
        var v37 uint64 = v2
        _ = v37
    case 42:
        var v38 float64 = v22
        _ = v38
    default:
        if len(v3) > 0 {
            v3[0] = 737
        }
    }
    if (float64(1.5) != (((v22 * float64(float32(v22))) / ((v22 + (-2.25 - ((1.5 + v22) - v22))) / 1.0)) + 1.5)) {
        switch "foo bar" {
        case "hello":
            if false {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                var v39 int = 172
                v5 -= -128
                _ = v39
            }
        case "foo bar":
            v40 := GenS4{F0: v5}
            _ = v40
        }
        v20 = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
