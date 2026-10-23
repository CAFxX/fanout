package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 uint16; F2 int16 }
type GenS5 struct { F0 int8; F1 uint }
func genFn0(p0 int32, p1 int64, p2 int64) uint64 {
    switch ((("" + ((("" + "") + "\n\t") + ("\n\t" + "foo bar"))) + ("foo bar" + "foo bar")) + ("foo bar" + "")) {
    case "\n\t":
        p2 = (p1 >> ((328) & 63))
    }
    switch 127 {
    case 42:
        p0 = p0
    case 0:
        switch ("\n\t" + "\n\t") {
        case "":
            v7 := GenS1{F0: 169, F1: 18446744073709551615, F2: 128}
            _ = v7
        case "a":
            p1 += p2
        case "\n\t":
            v8 := GenS1{F0: 1, F1: 2, F2: 32767}
            _ = v8
        }
    }
    if (("foo bar" + "hello") >= "\n\t") {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS2{F0: 110, F1: 42}
            _ = fi9
            _ = v10
        }
    } else {
        if ((!false) && (!(!(!((42 != 878) && true))))) {
            v11 := GenS2{F0: 38, F1: 216}
            for fi12 := 0; fi12 < 2; fi12++ {
                switch "\n\t" {
                case "foo bar":
                    for fi13 := 0; fi13 < 5; fi13++ {
                        fmt.Println(p0, fi12, p0)
                        var v14 float64 = float64(-2.25)
                        fmt.Println(p0)
                        _ = fi13
                        _ = v14
                    }
                }
                _ = fi12
            }
            var v15 uint32 = 4294967295
            _ = v11
            _ = v15
        } else {
            p0 += 0
            v16 := []int{212, 1, -2147483648, -602}
            _ = v16
        }
        p0 = (p0 << ((688) & 63))
        v17 := GenS3{F0: 67}
        _ = v17
    }
    if false {
        v18 := GenS2{F0: 127, F1: 255}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 uint64, p1 int64, p2 uint64) int16 {
    var v19 uint = 121
    fmt.Println(v19, v19, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 42
}

func main() {
    v1 := []string{"hello"}
    var v2 int8 = -128
    var v3 uint8 = 2
    v4 := []string{((("\n\t" + ("a" + "foo bar")) + "") + (((((("foo bar" + ("a" + "foo bar")) + ((("\n\t" + "hello") + "hello") + ("hello" + ""))) + "") + "\n\t") + "\n\t") + "")), ("\n\t" + "hello"), "a", ("hello" + (("\n\t" + ("\n\t" + "a")) + "foo bar"))}
    var v5 float32 = -2.25
    v6 := []uint32{1, 254}
    v5 = float32(3.14159)
    for fi20 := 0; fi20 < 3; fi20++ {
        if (-917 == 402) {
            var v21 int16 = genFn1(4294967295, 781, 456)
            switch "foo bar" {
            case "foo bar":
                var v22 string = (("foo bar" + ("\n\t" + "")) + ("foo bar" + "\n\t"))
                _ = v22
            default:
                v23 := GenS3{F0: 127}
                _ = v23
            }
            var v24 int8 = (v2 % (((v2 & 127)) | 1))
            _ = v21
            _ = v24
        }
        for fi25 := 0; fi25 < 16; fi25++ {
            v26 := GenS4{F0: (!(!((!false) || false))), F1: 33, F2: genFn1(4294967295, 0, 0)}
            _ = fi25
            _ = v26
        }
        _ = fi20
    }
    switch "\n\t" {
    case "":
        v6 = append(v6, 541)
    case "\n\t":
        v27 := GenS5{F0: int8(v3), F1: 450}
        _ = v27
    }
    fmt.Println(v5, v2)
    var v28 uint64 = 51
    var v29 float32 = ((1e10 * 0.0) / 3.14159)
    var v30 uint32 = 4294967295
    if false {
        fmt.Println(v2)
    }
    v29 = ((v29 + 0.0) + (v29 + v29))
    for fi31 := 0; fi31 < 19; fi31++ {
        fmt.Println(v28, v2)
        v32 := GenS1{F0: 18446744073709551615, F1: genFn0(-2147483648, 0, 1), F2: -9223372036854775808}
        var v33 int16 = 32767
        _ = fi31
        _ = v32
        _ = v33
    }
    v28 = (v28 >> ((255) & 63))
    for fi34 := 0; fi34 < 3; fi34++ {
        fmt.Println(v29, v3, v3)
        v35 := GenS3{F0: v3}
        var v36 float32 = 0.0
        _ = fi34
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        v4 = append(v4, "foo bar")
        _ = fi37
    }
    _ = len(v6)
    v2 += int8(v28)
    var v38 int16 = genFn1(v28, -1, v28)
    if (true && false) {
        fmt.Println(v38, v28)
        fmt.Println(v28, v5)
    }
    v29 = (v5 * v29)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
