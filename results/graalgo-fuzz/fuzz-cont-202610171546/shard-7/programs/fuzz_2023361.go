package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }

func main() {
    var v1 int8 = 127
    v2 := []string{("" + "\n\t")}
    var v3 int = 32767
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 int32 = 2147483647
        _ = fi4
        _ = v5
    }
    v1 -= v1
    var v6 int16 = -32768
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: -1, F1: "foo bar", F2: -1}
        v1 = (v1 * 100)
        var v9 float32 = (0.1 * (0.1 - float32(float64((1.5 * ((0.1 - 1.5) + (float32(float64(0.0)) * (0.0 / 1e10))))))))
        _ = fi7
        _ = v8
        _ = v9
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v6 = 0
    if ((v3 ^ 128) == v3) {
        fmt.Println(v3, v1)
    } else {
        if false {
            fmt.Println(v3, v6, v3)
            v6 -= v6
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: -551, F1: "", F2: -1}
            v12 := GenS1{F0: -1, F1: "a", F2: 0}
            _ = fi10
            _ = v11
            _ = v12
        }
        if len(v2) > 0 {
            v2[0] = ""
        }
    }
    var v13 int16 = v6
    var v14 uint = 0
    switch "\n\t" {
    case "foo bar":
        var v15 string = ("a" + "foo bar")
        _ = v15
    case "hello":
        for fi16 := 0; fi16 < 1; fi16++ {
            v14 = v14
            switch "foo bar" {
            case "":
                v14 = (v14 & v14)
            case "hello":
                v6 = (v13 - (v6 & v6))
            case "a":
                v2 = append(v2, "a")
            default:
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 bool = (!false)
                    v6 = v6
                    switch "\n\t" {
                    case "foo bar":
                        var v19 int = v3
                        _ = v19
                    case "\n\t":
                        v2 = append(v2, "hello")
                    case "hello":
                        var v20 uint = 0
                        if v18 {
                            v18 = (42 == 42)
                        }
                        switch (v3 - 547) {
                        case 42:
                            var v21 uint64 = 2
                            v14 -= v14
                            var v22 int16 = v13
                            _ = v21
                            _ = v22
                        }
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                }
            }
            if false {
                var v23 bool = (true && (0 == 2147483647))
                fmt.Println(v14, v3)
                if (-1 < 662) {
                    var v24 string = "a"
                    _ = len(v2)
                    _ = v24
                } else {
                    switch "a" {
                    case "":
                        _ = len(v2)
                    }
                }
                _ = v23
            }
            _ = fi16
        }
    }
    switch ("hello" + "foo bar") {
    case "hello":
        v14 += v14
    case "\n\t":
        if false {
            v1 = v1
            if (-237 != 931) {
                if len(v2) > 0 {
                    v2[0] = ""
                }
            } else {
                var v25 uint = (v14 / (((^(v14))) | 1))
                var v26 uint8 = 191
                if false {
                    v27 := GenS1{F0: 2147483647, F1: ("hello" + ("foo bar" + ("hello" + ("a" + "\n\t")))), F2: (v1 + 42)}
                    var v28 uint16 = 65535
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v27
                    _ = v28
                } else {
                    _ = len(v2)
                    v14 -= 397
                }
                _ = v25
                _ = v26
            }
        } else {
            v3 = (v3 | (v3 >> (((v14 & v14)) & 63)))
        }
    case "":
        var v29 uint16 = 193
        _ = v29
    }
    if ((^(v6)) < v6) {
        var v30 uint = (v14 >> ((v14) & 63))
        v13 -= v13
        var v31 int8 = (v1 / ((-128) | 1))
        _ = v30
        _ = v31
    }
    if true {
        v13 += (^(v13))
        switch "foo bar" {
        case "":
            _ = len(v2)
        }
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        _ = len(v2)
        _ = fi32
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        var v34 int32 = 2147483647
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        v6 -= (-(v6))
        var v36 int16 = v13
        v37 := GenS2{F0: 0, F1: (v3 % ((2147483647) | 1)), F2: (v1 & v1)}
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 uint = (v14 | v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v6
}
