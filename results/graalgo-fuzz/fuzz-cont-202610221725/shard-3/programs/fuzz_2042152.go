package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]string)
    var v3 uint64 = 4294967295
    var v4 uint = 1
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int64 = -583
        var v7 string = ""
        var v8 float32 = v1
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    if false {
        if (!(2 != 0)) {
            v1 = v1
            switch 1 {
            case 673:
                var v9 bool = false
                _ = v9
            default:
                switch 42 {
                case -1:
                    for fi10 := 0; fi10 < 13; fi10++ {
                        var v11 float32 = (v1 * 3.14159)
                        _ = fi10
                        _ = v11
                    }
                }
            }
            fmt.Println(v3)
        }
        var v12 uint64 = 614
        for fi13 := 0; fi13 < 19; fi13++ {
            fmt.Println(v1, v1)
            switch -272 {
            case -9223372036854775808:
                var v14 uint = 662
                _ = v14
            default:
                v3 = v12
            }
            switch 4294967295 {
            case -9223372036854775808:
                v15 := GenS1{F0: true, F1: (fi13 * 32768), F2: 100}
                _ = v15
            case -663:
                fmt.Println(fi13, v3, v4)
            }
            _ = fi13
        }
        _ = v12
    } else {
        if ("foo bar" >= ("\n\t" + "\n\t")) {
            for fi16 := 0; fi16 < 17; fi16++ {
                v3 = (v3 - v3)
                var v17 uint32 = 2
                _ = fi16
                _ = v17
            }
        } else {
            var v18 float64 = 1.5
            v18 -= v18
            var v19 uint8 = 127
            _ = v18
            _ = v19
        }
    }
    switch -382 {
    case 32768:
        v2[65535] = "foo bar"
    }
    if ((false || (1 <= uint8(v4))) && false) {
        var v20 float32 = ((3.14159 * (0.0 / 1.0)) + v1)
        _ = v20
    }
    v1 = 1e10
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        for fi21 := 0; fi21 < 3; fi21++ {
            v3 += v3
            var v22 uint = (^(v4))
            _ = len(v2)
            _ = fi21
            _ = v22
        }
    }
    fmt.Println(v1)
    v23 := []int{948, -9223372036854775808, 9223372036854775807, -9223372036854775808}
    v24 := GenS1{F0: ("a" >= ("hello" + "hello")), F1: 42, F2: 0}
    for fi25 := 0; fi25 < 7; fi25++ {
        v26 := GenS2{F0: 110}
        _ = fi25
        _ = v26
    }
    if (1 >= 42) {
        v27 := GenS3{F0: "hello"}
        var v28 uint16 = 277
        switch 9223372036854775807 {
        case 4294967295:
            var v29 int64 = -1
            _ = v29
        case 42:
            _ = len(v2)
        }
        _ = v27
        _ = v28
    } else {
        switch ("" + ("hello" + "foo bar")) {
        case "a":
            var v30 int8 = 1
            _ = v30
        case "":
            for fi31 := 0; fi31 < 19; fi31++ {
                var v32 int = fi31
                _ = fi31
                _ = v32
            }
        default:
            if (true && true) {
                v33 := GenS2{F0: 123}
                var v34 int = -2147483648
                _ = v33
                _ = v34
            } else {
                _ = len(v23)
            }
        }
        fmt.Println(v1, v3)
    }
    if true {
        for fi35 := 0; fi35 < 3; fi35++ {
            v36 := v2[1]
            switch (fi35 << (((v4 & (v4 + (v4 >> ((v4) & 63))))) & 63)) {
            case 9223372036854775807:
                if (!true) {
                    var v37 uint8 = 255
                    v2[2] = v36
                    if true {
                        var v38 float32 = ((1e10 - ((0.0 * (((v1 / v1) / ((0.1 - 1.5) / v1)) + (v1 * v1))) + v1)) + (v1 * 3.14159))
                        v1 = 3.14159
                        v2[2] = (((((((("a" + v36) + ("foo bar" + v36)) + (v36 + v36)) + "\n\t") + v36) + "a") + ("hello" + v36)) + "")
                        _ = v38
                    } else {
                        v3 += v3
                        fmt.Println(v36, v3)
                    }
                    _ = v37
                }
            }
            _ = fi35
            _ = v36
        }
        _ = len(v2)
        v3 += v3
    } else {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        v3 = v3
    }
    var v39 uint32 = 2
    var v40 float64 = (1e10 + 1.5)
    var v41 float64 = 0.0
    if false {
        v41 += v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
