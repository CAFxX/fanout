package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 bool }
type GenS2 struct { F0 int; F1 int32; F2 uint8 }
type GenS3 struct { F0 int; F1 uint64; F2 bool }
type GenS4 struct { F0 bool; F1 bool; F2 uint16 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int32) int {
    var v5 uint = 255
    v5 = v5
    if false {
        var v6 uint64 = 262
        fmt.Println(v6)
        _ = v6
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v7 := []int{}
        _ = v7
    case "hello":
        v5 = 371
    case "a":
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = v5
    return int(p0)
}
func genFn1() int16 {
    var v8 string = ""
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint64 = 18446744073709551615
        v11 := GenS1{F0: (fi9 - (fi9 - (-(fi9)))), F1: 42, F2: ((false || false) || (!(-776 < -75)))}
        v12 := []int{fi9, 320, fi9}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{2147483647, 32768, -1}
    _ = v13
    _ = v8
    return 0
}

func main() {
    v1 := []string{"foo bar", ""}
    var v2 float64 = (3.14159 + 1.5)
    v3 := []string{(("" + "") + "foo bar"), ("foo bar" + "foo bar"), ((("" + ("foo bar" + "foo bar")) + "") + "a")}
    v4 := make(map[uint32]bool)
    for fi14 := 0; fi14 < 12; fi14++ {
        if true {
            v2 += v2
        } else {
            v2 = v2
            fmt.Println(fi14, fi14)
        }
        switch "foo bar" {
        case "a":
            v2 = v2
        case "hello":
            v15 := GenS2{F0: fi14, F1: 1, F2: 0}
            _ = v15
        }
        _ = fi14
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    var v16 int32 = -482
    if false {
        fmt.Println(v16, v2)
        if (32767 < 0) {
            for fi17 := 0; fi17 < 12; fi17++ {
                for fi18 := 0; fi18 < 2; fi18++ {
                    fmt.Println(fi17)
                    if ("\n\t" < "\n\t") {
                        v19 := GenS3{F0: fi17, F1: 2, F2: (0 >= -1)}
                        if (!true) {
                            var v20 int16 = genFn1()
                            _ = v20
                        }
                        fmt.Println(fi17, v16)
                        _ = v19
                    }
                    v16 = -625
                    _ = fi18
                }
                v2 = v2
                _ = fi17
            }
            v21 := GenS3{F0: -246, F1: 924, F2: true}
            _ = v21
        } else {
            v16 = (v16 / (((v16 * 0)) | 1))
            v16 += v16
            v22 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: (!false)}
            _ = v22
        }
    } else {
        v23 := GenS2{F0: -2147483648, F1: v16, F2: 151}
        for fi24 := 0; fi24 < 15; fi24++ {
            switch "\n\t" {
            case "foo bar":
                v25 := GenS3{F0: 423, F1: 18446744073709551615, F2: ((1 > 1) || (v2 <= 0.0))}
                _ = v25
            case "a":
                var v26 uint64 = 1
                _ = v26
            }
            var v27 int = 41
            _ = fi24
            _ = v27
        }
        _ = v23
    }
    var v28 float32 = 1e10
    fmt.Println(v2)
    v29 := GenS4{F0: (v2 != v2), F1: false, F2: 655}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 uint8 = 217
        v32 := GenS5{F0: 670}
        v28 += v28
        _ = fi30
        _ = v31
        _ = v32
    }
    v16 = (v16 * v16)
    fmt.Println(v2)
    fmt.Println(v16)
    if ((true || false) || (!(!(true || (true && (2147483647 > genFn0(0))))))) {
        for fi33 := 0; fi33 < 20; fi33++ {
            if false {
                v2 += (1e10 + v2)
            } else {
                var v34 int8 = -128
                switch "foo bar" {
                case "":
                    var v35 uint16 = 2
                    _ = v35
                }
                fmt.Println(v34)
                _ = v34
            }
            fmt.Println(v28, fi33, v2)
            var v36 uint16 = 32767
            _ = fi33
            _ = v36
        }
    }
    var v37 float32 = 0.0
    var v38 uint = 18446744073709551615
    for fi39 := 0; fi39 < 20; fi39++ {
        fmt.Println(v37, fi39, v38)
        for fi40 := 0; fi40 < 15; fi40++ {
            if false {
                var v41 uint16 = 2
                _ = len(v4)
                v42 := GenS4{F0: (uint32(v38) >= 784), F1: (!(!(0.1 == (v2 - v2)))), F2: v41}
                _ = v41
                _ = v42
            }
            v28 = float32(v2)
            var v43 int32 = -642
            _ = fi40
            _ = v43
        }
        _ = fi39
    }
    if false {
        var v44 int64 = 0
        _ = v44
    }
    var v45 int8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v45
}
