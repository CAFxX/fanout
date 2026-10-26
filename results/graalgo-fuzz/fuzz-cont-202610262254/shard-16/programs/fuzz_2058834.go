package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint64; F1 bool; F2 uint16 }

func main() {
    var v1 string = ((((("hello" + (("\n\t" + ("" + "hello")) + "")) + "foo bar") + (("\n\t" + "\n\t") + "hello")) + ((("hello" + ("" + "\n\t")) + "\n\t") + "\n\t")) + ("\n\t" + ("" + (((("" + "foo bar") + "") + "") + "hello"))))
    v2 := []int{0}
    var v3 uint8 = 0
    v4 := []uint32{65535, 65535}
    v5 := make(map[string]bool)
    v1 = ""
    v6 := GenS1{F0: 2}
    v1 = ("a" + ((v1 + ((v1 + (("\n\t" + v1) + (v1 + ("" + v1)))) + (((("" + v1) + v1) + "\n\t") + (("\n\t" + "foo bar") + v1)))) + (v1 + ("hello" + v1))))
    fmt.Println(v1, v3)
    var v7 int8 = 127
    var v8 int16 = -630
    if (664 > 636) {
        switch -9223372036854775808 {
        case 0:
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := GenS2{F0: "\n\t"}
                var v11 int64 = -1
                _ = fi9
                _ = v10
                _ = v11
            }
        case 4294967295:
            var v12 int = 127
            _ = v12
        case -1:
            var v13 int64 = 42
            _ = v13
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int16 = 837
            if (!true) {
                switch fi14 {
                case 2147483647:
                    v8 = v15
                case 128:
                    v15 = (v8 | (v15 / ((v8) | 1)))
                }
                v15 += 0
                v5[((v1 + "foo bar") + (v1 + ((("hello" + v1) + v1) + "")))] = (1 > -1)
            }
            _ = fi14
            _ = v15
        }
        switch 32767 {
        case 32767:
            var v16 int16 = (v8 << ((65535) & 63))
            _ = v16
        case 42:
            for fi17 := 0; fi17 < 13; fi17++ {
                for fi18 := 0; fi18 < 17; fi18++ {
                    v3 = (v3 / (((v3 << ((1) & 63))) | 1))
                    v19 := GenS3{F0: 853, F1: 0, F2: (v3 % ((102) | 1))}
                    v8 -= -651
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
        }
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        v1 = v1
        if true {
            v8 = (v8 / (((v8 + int16(v3))) | 1))
            var v21 uint32 = 2
            _ = v21
        } else {
            var v22 int64 = 622
            _ = v22
        }
        v7 -= v7
        _ = fi20
    }
    if false {
        var v23 uint8 = (v3 & 61)
        var v24 uint64 = 4294967295
        _ = v23
        _ = v24
    }
    var v25 float32 = 1e10
    if false {
        var v26 int = 32767
        if ((1e10 / (0.1 + (1.5 * 0.1))) < float64(float32((0.0 + (1.5 / (float64((v25 + v25)) * 1.5)))))) {
            for fi27 := 0; fi27 < 20; fi27++ {
                if (780 >= 255) {
                    var v28 uint64 = 225
                    var v29 uint16 = 600
                    _ = v28
                    _ = v29
                }
                var v30 int64 = -1
                _ = fi27
                _ = v30
            }
        } else {
            switch (v1 + v1) {
            case "":
                var v31 uint8 = (v3 << ((60) & 63))
                _ = v31
            }
        }
        _ = v26
    } else {
        var v32 uint32 = 0
        v7 -= v7
        for fi33 := 0; fi33 < 16; fi33++ {
            v34 := v5[(v1 + (((v1 + "") + v1) + "a"))]
            _ = fi33
            _ = v34
        }
        _ = v32
    }
    if (false || (1 >= v7)) {
        v7 = (v7 & v7)
        var v35 int16 = v8
        var v36 int64 = 1
        _ = v35
        _ = v36
    } else {
        v37 := GenS4{F0: 18446744073709551615, F1: ((v3 << ((255) & 63)) >= v3), F2: 0}
        v38 := v5[v1]
        _ = v37
        _ = v38
    }
    fmt.Println(v25, v3)
    v7 = v7
    var v39 int32 = -1
    switch v1 {
    case "a":
        var v40 uint = 255
        _ = v40
    case "\n\t":
        var v41 int32 = (v39 ^ (v39 % (((v39 * v39)) | 1)))
        _ = v41
    }
    fmt.Println(v7, v8, v1)
    fmt.Println(v25, v7, v39)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
