package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }

func main() {
    v1 := []int{32767}
    var v2 uint64 = 1
    var v3 bool = (2 != 42)
    var v4 uint16 = 255
    for fi5 := 0; fi5 < 4; fi5++ {
        _ = len(v1)
        _ = fi5
    }
    if ("a" > ("a" + "foo bar")) {
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 int = (fi6 - 42)
            _ = fi6
            _ = v7
        }
        var v8 string = ("a" + "")
        v4 = (v4 / (((v4 % ((v4) | 1))) | 1))
        _ = v8
    } else {
        var v9 bool = (!v3)
        fmt.Println(v3, v2, v4)
        _ = v9
    }
    if len(v1) > 0 {
        v1[0] = 102
    }
    v2 = v2
    if true {
        var v10 int8 = 0
        v11 := GenS1{F0: -1, F1: 1, F2: 420}
        v2 = (^(v2))
        _ = v10
        _ = v11
    } else {
        v3 = v3
        var v12 uint16 = 0
        for fi13 := 0; fi13 < 20; fi13++ {
            if true {
                var v14 uint32 = 42
                _ = v14
            } else {
                var v15 int32 = 42
                for fi16 := 0; fi16 < 12; fi16++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v12 = 32767
                    _ = fi16
                }
                _ = v15
            }
            v12 += v12
            _ = fi13
        }
        _ = v12
    }
    if (1000 != -1) {
        switch ("" + "a") {
        case "\n\t":
            for fi17 := 0; fi17 < 8; fi17++ {
                if v3 {
                    var v18 uint16 = v4
                    for fi19 := 0; fi19 < 9; fi19++ {
                        fmt.Println(fi19)
                        if false {
                            var v20 int16 = -1
                            _ = v20
                        }
                        _ = fi19
                    }
                    v3 = ((v18 < (v18 ^ (v4 * v4))) || (v3 || (0 < 9223372036854775807)))
                    _ = v18
                }
                var v21 string = (("foo bar" + "") + "a")
                _ = fi17
                _ = v21
            }
        case "":
            v1 = append(v1, -80)
        case "foo bar":
            _ = len(v1)
        }
        fmt.Println(v4)
        v1 = append(v1, -841)
    }
    var v22 bool = (uint16(v2) <= 2)
    var v23 int = -1
    for fi24 := 0; fi24 < 1; fi24++ {
        v23 = fi24
        _ = fi24
    }
    v4 -= v4
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS1{F0: 1, F1: 999, F2: 278}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 = ((fi25 | -2147483648) == -2147483648)
        _ = fi25
        _ = v26
    }
    var v27 int8 = 127
    for fi28 := 0; fi28 < 3; fi28++ {
        fmt.Println(v3)
        var v29 uint16 = (v4 + (v4 + 42))
        for fi30 := 0; fi30 < 1; fi30++ {
            fmt.Println(v23, v29, v3)
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    if false {
        v23 -= v23
        if len(v1) > 0 {
            v1[0] = (v23 + v23)
        }
        v22 = ((v27 / ((v27) | 1)) > (v27 + (v27 + (v27 & (-(v27))))))
    }
    fmt.Println(v22, v23)
    v31 := GenS2{F0: (v2 | v2)}
    if (v22 || true) {
        var v32 uint16 = 2
        _ = v32
    } else {
        switch (v23 % (((v23 & -948)) | 1)) {
        case 9223372036854775807:
            for fi33 := 0; fi33 < 2; fi33++ {
                switch "" {
                case "":
                    switch (v23 & (-(fi33))) {
                    case 483:
                        var v34 uint32 = 1
                        _ = v34
                    case -9223372036854775808:
                        v3 = (v22 || (0.1 > 0.1))
                    case 1:
                        fmt.Println(v3, v22)
                    }
                case "a":
                    if ("\n\t" <= "a") {
                        if (-423 > -866) {
                            v35 := GenS3{F0: int64(v23)}
                            v36 := GenS2{F0: v2}
                            _ = v35
                            _ = v36
                        }
                        var v37 bool = ((v4 * 0) == v4)
                        _ = v37
                    }
                }
                var v38 int32 = 2147483647
                if (0.0 == (0.0 - float32(float64(3.14159)))) {
                    v39 := GenS1{F0: int64(v38), F1: (v38 << ((603) & 63)), F2: 2}
                    v40 := GenS4{F0: 65535}
                    _ = v39
                    _ = v40
                } else {
                    v41 := GenS1{F0: 1, F1: 0, F2: 1}
                    if v22 {
                        var v42 int16 = 32767
                        fmt.Println(v38, v27, v22)
                        _ = v42
                    }
                    var v43 float32 = 0.0
                    _ = v41
                    _ = v43
                }
                _ = fi33
                _ = v38
            }
        case -2147483648:
            v44 := GenS3{F0: 157}
            _ = v44
        }
        var v45 string = "foo bar"
        _ = v45
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        var v47 string = ("hello" + (("\n\t" + ("hello" + "\n\t")) + "foo bar"))
        _ = fi46
        _ = v47
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v31
    _ = v4
}
