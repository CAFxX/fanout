package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint8 }

func main() {
    var v1 float32 = (-2.25 - (float32(-2.25) + float32(1.5)))
    var v2 uint = 0
    v3 := []string{"", (("a" + ("\n\t" + "foo bar")) + "foo bar")}
    var v4 uint32 = 1
    var v5 float64 = (1e10 * 0.1)
    v6 := make(map[string]string)
    v1 += v1
    v4 -= (v4 << ((2) & 63))
    if false {
        fmt.Println(v2, v1)
        if false {
            if false {
                for fi7 := 0; fi7 < 2; fi7++ {
                    var v8 int32 = -721
                    _ = fi7
                    _ = v8
                }
                if len(v3) > 0 {
                    v3[0] = ("" + "\n\t")
                }
            } else {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 int16 = -1
                    v11 := GenS1{F0: (fi9 >> (((v2 >> (((v2 << ((18446744073709551615) & 63))) & 63))) & 63)), F1: int64(v10)}
                    v1 += v1
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        }
        v4 = (v4 % (((v4 ^ v4)) | 1))
    } else {
        var v12 string = ""
        switch v12 {
        case "foo bar":
            v1 -= 1.5
        case "a":
            switch ((v12 + (v12 + "hello")) + "foo bar") {
            case "hello":
                var v13 bool = true
                _ = v13
            case "a":
                switch (v12 + ((v12 + v12) + ((v12 + "hello") + "a"))) {
                case "":
                    v14 := v6["a"]
                    _ = v14
                }
            default:
                v1 = ((v1 + float32(3.14159)) - (1.5 * v1))
            }
        case "":
            delete(v6, (("foo bar" + v12) + "a"))
        }
        switch "foo bar" {
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = ("hello" + v12)
            }
        }
        _ = v12
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = v2
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS1{F0: (fi15 & 32767), F1: 861}
        if true {
            switch (("" + "\n\t") + ((("" + "a") + "") + "foo bar")) {
            case "a":
                delete(v6, "")
            case "\n\t":
                v4 -= v4
            }
            if true {
                _ = len(v6)
            }
        }
        fmt.Println(v1, v2)
        _ = fi15
        _ = v16
    }
    switch -2147483648 {
    case 2147483647:
        if false {
            v4 = v4
        }
    case -1:
        v5 += (float64((((v1 - 0.1) + v1) * (v1 + float32(0.1)))) - (v5 * (-2.25 * float64(0.1))))
    }
    v17 := GenS1{F0: -658, F1: -1}
    for fi18 := 0; fi18 < 10; fi18++ {
        switch "" {
        case "":
            if (!true) {
                if len(v3) > 0 {
                    v3[0] = ("foo bar" + "\n\t")
                }
            }
        case "a":
            var v19 uint = (v2 % ((255) | 1))
            _ = v19
        }
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := GenS1{F0: fi20, F1: -1}
            _ = fi20
            _ = v21
        }
        _ = fi18
    }
    fmt.Println(v5, v2)
    if (("hello" > ("foo bar" + "a")) || (1 >= 42)) {
        if true {
            _ = len(v6)
            var v22 float32 = v1
            v1 += v1
            _ = v22
        } else {
            _ = len(v3)
            for fi23 := 0; fi23 < 7; fi23++ {
                if true {
                    v24 := v6["a"]
                    _ = v24
                } else {
                    v4 += v4
                    v1 = 3.14159
                }
                _ = fi23
            }
        }
        v25 := GenS2{F0: 2147483647}
        switch -251 {
        case 2147483647:
            fmt.Println(v5)
        case -452:
            v26 := GenS2{F0: 2147483647}
            _ = v26
        case -2147483648:
            var v27 float32 = v1
            _ = v27
        }
        _ = v25
    } else {
        v1 = v1
    }
    switch 0 {
    case 685:
        for fi28 := 0; fi28 < 6; fi28++ {
            v29 := GenS3{F0: 127, F1: 255}
            _ = len(v3)
            _ = fi28
            _ = v29
        }
    case 1:
        v30 := GenS1{F0: 4294967295, F1: 42}
        _ = v30
    default:
        v1 = ((((0.0 + (v1 / v1)) * v1) + 0.0) + v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
