package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 string; F2 string }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    var v3 string = ("foo bar" + "a")
    var v4 int32 = -301
    var v5 float64 = float64(0.0)
    v3 = ""
    for fi6 := 0; fi6 < 19; fi6++ {
        switch fi6 {
        case 0:
            v7 := v1[(fi6 & fi6)]
            _ = v7
        default:
            var v8 int32 = v4
            _ = v8
        }
        if (!((((v3 + v3) > v3) || false) && true)) {
            v4 = v4
            v5 -= (0.0 / -2.25)
        }
        v9 := v1[-358]
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        switch (v3 + ("foo bar" + "")) {
        case "":
            var v11 bool = true
            _ = v11
        case "a":
            v12 := GenS1{F0: false, F1: 0}
            _ = v12
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 string = ("foo bar" + v3)
        for fi15 := 0; fi15 < 7; fi15++ {
            switch 615 {
            case -59:
                v16 := GenS1{F0: true, F1: 17}
                if (1 >= 42) {
                    fmt.Println(fi13, v5, fi13)
                } else {
                    delete(v2, (fi15 << ((255) & 63)))
                }
                _ = v16
            case -223:
                var v17 int16 = -32768
                _ = v17
            case -2147483648:
                v4 = 1
            }
            v18 := GenS2{F0: 4294967295, F1: 2147483647}
            _ = fi15
            _ = v18
        }
        _ = fi13
        _ = v14
    }
    v4 -= (v4 << ((65535) & 63))
    fmt.Println(v4, v4)
    v19 := GenS3{F0: 42, F1: 1, F2: 1}
    for fi20 := 0; fi20 < 8; fi20++ {
        switch 1 {
        case -2147483648:
            if (32767 == 2) {
                for fi21 := 0; fi21 < 3; fi21++ {
                    var v22 uint16 = 32767
                    _ = fi21
                    _ = v22
                }
            } else {
                if (!true) {
                    fmt.Println(v5, fi20)
                } else {
                    fmt.Println(v5, fi20)
                }
            }
        case 128:
            v4 = 488
        }
        var v23 int16 = -902
        _ = fi20
        _ = v23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        v25 := []int{-2147483648}
        switch ("" + "") {
        case "foo bar":
            v5 = (v5 - (v5 / 1.5))
        case "a":
            for fi26 := 0; fi26 < 14; fi26++ {
                if (!false) {
                    switch v3 {
                    case "\n\t":
                        v5 = -2.25
                    case "a":
                        v27 := GenS4{F0: -128, F1: 0}
                        _ = v27
                    case "":
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                    }
                    v3 = (("foo bar" + ((v3 + v3) + "hello")) + ((((("" + v3) + v3) + ("hello" + ("\n\t" + v3))) + "foo bar") + "foo bar"))
                    var v28 uint64 = 18446744073709551615
                    _ = v28
                }
                _ = fi26
            }
        case "hello":
            v3 = v3
        }
        _ = fi24
        _ = v25
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := GenS5{F0: v3}
        _ = fi29
        _ = v30
    }
    var v31 string = ("foo bar" + "hello")
    switch -1 {
    case 0:
        v32 := GenS5{F0: (("foo bar" + (("a" + v31) + (v31 + v31))) + v3)}
        _ = v32
    default:
        v31 = "\n\t"
    }
    v33 := GenS6{F0: v4, F1: v3, F2: v31}
    var v34 string = "\n\t"
    var v35 int64 = -9223372036854775808
    v3 = (((("\n\t" + "") + ((v31 + v31) + v31)) + v31) + ((("a" + (v34 + "")) + "") + v34))
    var v36 uint16 = 68
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
