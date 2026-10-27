package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint = 65535
    var v2 uint64 = 42
    v3 := make(map[uint32]int)
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 12; fi5++ {
        v2 = 0
        _ = fi5
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v6 int8 = -1
        _ = v6
    case 32768:
        v7 := []int{2147483647}
        _ = v7
    default:
        if false {
            var v8 uint = v1
            if (!((1e10 * 1e10) >= 0.1)) {
                var v9 uint32 = 224
                if true {
                    var v10 float32 = 0.1
                    _ = v10
                }
                _ = v9
            }
            v2 = 2
            _ = v8
        } else {
            switch "" {
            case "a":
                switch ("a" + "foo bar") {
                case "\n\t":
                    for fi11 := 0; fi11 < 11; fi11++ {
                        fmt.Println(fi11, v2, fi11)
                        var v12 float32 = 3.14159
                        v2 -= 678
                        _ = fi11
                        _ = v12
                    }
                }
            }
            v4[1] = "hello"
        }
    }
    v13 := []int{0}
    v14 := GenS1{F0: false}
    _ = len(v3)
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS1{F0: (fi15 > (fi15 ^ (fi15 / ((fi15) | 1))))}
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: (true && false)}
    if true {
        switch 32767 {
        case 1:
            if false {
                v18 := GenS2{F0: "\n\t"}
                fmt.Println(v2, v1)
                _ = v18
            } else {
                if (-759 != -1) {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        fmt.Println(fi19, fi19, v2)
                        fmt.Println(v1, fi19)
                        var v20 uint32 = 0
                        _ = fi19
                        _ = v20
                    }
                    var v21 float32 = 0.0
                    var v22 uint32 = 2147483648
                    _ = v21
                    _ = v22
                } else {
                    var v23 int32 = 0
                    _ = v23
                }
                _ = len(v3)
                fmt.Println(v2)
            }
        case -9223372036854775808:
            var v24 int32 = -924
            _ = v24
        }
        for fi25 := 0; fi25 < 19; fi25++ {
            for fi26 := 0; fi26 < 16; fi26++ {
                fmt.Println(fi26, v1, v2)
                _ = fi26
            }
            switch (fi25 >> ((v1) & 63)) {
            case -823:
                fmt.Println(v1, fi25)
            case -2147483648:
                for fi27 := 0; fi27 < 1; fi27++ {
                    if (65535 >= 255) {
                        switch (fi27 & fi25) {
                        case 42:
                            var v28 uint = 385
                            _ = v28
                        }
                    }
                    v29 := GenS3{F0: 230}
                    _ = fi27
                    _ = v29
                }
            }
            if (65535 == 2147483647) {
                var v30 uint32 = 32
                var v31 int8 = 0
                var v32 int16 = 42
                _ = v30
                _ = v31
                _ = v32
            } else {
                if (!((1.5 + 1.5) >= float64(float32((3.14159 + 0.1))))) {
                    for fi33 := 0; fi33 < 6; fi33++ {
                        fmt.Println(v1, fi33, v1)
                        delete(v4, fi25)
                        _ = fi33
                    }
                    var v34 bool = false
                    _ = v34
                } else {
                    fmt.Println(fi25, v1, v1)
                }
                v35 := GenS4{F0: 638}
                var v36 uint64 = v2
                _ = v35
                _ = v36
            }
            _ = fi25
        }
        var v37 float32 = 1.5
        _ = v37
    }
    v38 := GenS2{F0: (("" + "\n\t") + "a")}
    v39 := GenS2{F0: "foo bar"}
    if true {
        v1 = v1
        v40 := v3[65535]
        var v41 float32 = 1.5
        _ = v40
        _ = v41
    } else {
        v4[1] = (("a" + "hello") + ("a" + ""))
        var v42 bool = (3.14159 != 3.14159)
        _ = v42
    }
    var v43 uint = 935
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
}
