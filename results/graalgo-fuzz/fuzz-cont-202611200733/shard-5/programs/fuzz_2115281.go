package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := []string{"a", "a"}
    var v3 uint = 1
    switch 32768 {
    case 127:
        var v4 uint16 = 42
        _ = v4
    default:
        if ((-2.25 / -2.25) <= float32(3.14159)) {
            if (0.0 == 3.14159) {
                var v5 uint64 = 2
                _ = v5
            } else {
                v3 = v3
            }
        } else {
            fmt.Println(v3)
            v6 := GenS1{F0: "foo bar"}
            _ = v6
        }
    }
    var v7 bool = true
    var v8 float32 = (float32(1.5) / 1.0)
    if false {
        v2 = append(v2, ("\n\t" + "foo bar"))
        v3 += 255
        var v9 uint8 = 1
        _ = v9
    } else {
        var v10 bool = v7
        fmt.Println(v8, v8, v3)
        _ = v10
    }
    switch 32768 {
    case -2147483648:
        var v11 int64 = 1
        _ = v11
    case 727:
        var v12 int32 = 1
        _ = v12
    default:
        v1[(("hello" + ("hello" + "a")) + ("hello" + ("foo bar" + ("a" + ("a" + ("a" + ("hello" + ("hello" + "hello"))))))))] = -9223372036854775808
    }
    var v13 string = ""
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            fmt.Println(fi15, fi15)
            v16 := v1[v13]
            v17 := GenS2{F0: 1}
            _ = fi15
            _ = v16
            _ = v17
        }
        switch 1 {
        case 0:
            if true {
                _ = len(v2)
                if (!(!false)) {
                    switch (fi14 & -547) {
                    case 4294967295:
                        v8 -= v8
                    case 42:
                        fmt.Println(v13, fi14)
                    case 128:
                        for fi18 := 0; fi18 < 14; fi18++ {
                            var v19 uint64 = 1
                            v20 := GenS3{F0: v13, F1: 42}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    }
                    fmt.Println(v8, v7, v3)
                }
                if len(v2) > 0 {
                    v2[0] = v13
                }
            } else {
                for fi21 := 0; fi21 < 13; fi21++ {
                    var v22 uint16 = 119
                    switch (fi21 ^ fi14) {
                    case 2147483647:
                        v7 = v7
                    }
                    _ = fi21
                    _ = v22
                }
            }
        case -972:
            if len(v2) > 0 {
                v2[0] = v13
            }
        }
        _ = fi14
    }
    v3 += (v3 | v3)
    switch v13 {
    case "":
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS2{F0: 192}
            _ = fi23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 18; fi26++ {
            switch (fi25 << ((v3) & 63)) {
            case 32767:
                switch fi26 {
                case 32768:
                    switch (("a" + v13) + "foo bar") {
                    case "":
                        v13 = ""
                    }
                }
            }
            _ = fi26
        }
        _ = fi25
    }
    v27 := GenS1{F0: (("hello" + v13) + v13)}
    fmt.Println(v3, v13, v7)
    fmt.Println(v13, v13)
    var v28 string = v13
    _ = len(v1)
    var v29 int16 = 882
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v7
    _ = v8
}
