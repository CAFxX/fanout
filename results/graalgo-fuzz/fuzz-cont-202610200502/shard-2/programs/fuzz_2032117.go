package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 bool }
type GenS3 struct { F0 uint32; F1 string; F2 int32 }

func main() {
    v1 := []int{-284, -463, 127, 127}
    var v2 float64 = 0.0
    var v3 uint16 = 2
    var v4 string = ((("a" + "a") + ("hello" + "")) + "a")
    v5 := []string{v4, v4, "\n\t"}
    _ = len(v5)
    if false {
        v6 := GenS1{F0: 182}
        switch 332 {
        case -1:
            var v7 int16 = 32767
            _ = v7
        }
        _ = v6
    } else {
        switch 1 {
        case 32768:
            var v8 uint64 = 129
            _ = v8
        case 4294967295:
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 float32 = 1e10
                var v11 uint8 = 62
                v4 = ((v4 + ((v4 + (v4 + v4)) + "")) + v4)
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        _ = len(v5)
        v5 = append(v5, v4)
    }
    var v12 int64 = 9223372036854775807
    for fi13 := 0; fi13 < 6; fi13++ {
        switch fi13 {
        case -9223372036854775808:
            if len(v1) > 0 {
                v1[0] = 34
            }
        case 32767:
            switch ((v4 + (v4 + "")) + v4) {
            case "":
                fmt.Println(v3, v12)
            case "hello":
                var v14 uint64 = 42
                _ = v14
            }
        default:
            if false {
                fmt.Println(v3, v2)
                _ = len(v5)
            } else {
                var v15 uint8 = uint8(v3)
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            switch 128 {
            case 196:
                v5 = append(v5, (v4 + ((v4 + v4) + ((v4 + v4) + (("\n\t" + "a") + v4)))))
            case 1:
                if (!(!true)) {
                    var v17 float64 = v2
                    if len(v5) > 0 {
                        v5[0] = v4
                    }
                    var v18 uint16 = v3
                    _ = v17
                    _ = v18
                }
            default:
                v19 := GenS1{F0: 128}
                _ = v19
            }
            v3 = v3
            _ = fi16
        }
        _ = fi13
    }
    if (!(1 < 500)) {
        fmt.Println(v12, v2)
    }
    if false {
        if true {
            if true {
                var v20 int8 = 82
                v20 -= 100
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v20
            } else {
                _ = len(v1)
                v2 = v2
                v4 = (v4 + "foo bar")
            }
            switch (v4 + ((v4 + v4) + "foo bar")) {
            case "a":
                if len(v1) > 0 {
                    v1[0] = 42
                }
            default:
                var v21 float32 = (1.5 - 1.5)
                _ = v21
            }
            var v22 float64 = v2
            _ = v22
        } else {
            v23 := GenS2{F0: (v3 & v3), F1: 0, F2: (!((255 != 42) && true))}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v2, v4, v12)
            _ = v23
        }
    } else {
        v24 := GenS3{F0: 2, F1: v4, F2: 0}
        for fi25 := 0; fi25 < 6; fi25++ {
            _ = len(v1)
            _ = fi25
        }
        if ((v3 >> ((255) & 63)) >= (v3 - (v3 - (v3 * (v3 << ((255) & 63)))))) {
            fmt.Println(v12)
        } else {
            fmt.Println(v3, v2, v12)
        }
        _ = v24
    }
    switch "a" {
    case "":
        var v26 bool = false
        _ = v26
    case "foo bar":
        var v27 string = v4
        _ = v27
    case "a":
        if (true || ((^(v3)) > v3)) {
            v3 = 890
        } else {
            var v28 uint16 = (v3 & v3)
            if (!true) {
                v3 -= v3
                _ = len(v5)
            } else {
                fmt.Println(v12, v12)
            }
            _ = len(v5)
            _ = v28
        }
    }
    v29 := GenS3{F0: 42, F1: ((v4 + ((v4 + ("hello" + (v4 + ("foo bar" + "hello")))) + ("\n\t" + v4))) + (v4 + v4)), F2: 1}
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 uint16 = v3
        var v32 string = ((("\n\t" + v4) + (v4 + ("\n\t" + ("a" + v4)))) + ((v4 + v4) + "\n\t"))
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
