package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 string }

func main() {
    var v1 string = "foo bar"
    var v2 float32 = 1.5
    v3 := []string{v1, "\n\t", (("\n\t" + ((v1 + v1) + v1)) + v1), v1}
    var v4 uint32 = 2147483647
    var v5 int16 = 1
    v6 := GenS1{F0: 716, F1: 1}
    if (!false) {
        fmt.Println(v2, v4)
        if false {
            v3 = append(v3, ((v1 + v1) + v1))
            var v7 bool = true
            v8 := GenS1{F0: -2147483648, F1: 99}
            _ = v7
            _ = v8
        }
    } else {
        v9 := GenS1{F0: -2147483648, F1: -1}
        _ = v9
    }
    var v10 float32 = (float32(float64(0.0)) / float32(1e10))
    v10 += v2
    for fi11 := 0; fi11 < 11; fi11++ {
        v4 += 65535
        var v12 uint = 0
        switch (^(fi11)) {
        case 697:
            if (!true) {
                var v13 string = ("foo bar" + v1)
                for fi14 := 0; fi14 < 1; fi14++ {
                    v15 := GenS1{F0: -2147483648, F1: 127}
                    if len(v3) > 0 {
                        v3[0] = (v13 + ((v13 + (v1 + "foo bar")) + (v1 + v1)))
                    }
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
        case -71:
            switch ("a" + v1) {
            case "a":
                v16 := GenS1{F0: 42, F1: 100}
                _ = v16
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "hello":
                _ = len(v3)
            }
        case 127:
            v3 = append(v3, v1)
        default:
            v3 = append(v3, v1)
        }
        _ = fi11
        _ = v12
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v2 = float32(float64((v10 / v2)))
        if false {
            var v18 bool = false
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int64 = 9223372036854775807
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 int64 = -1
            fmt.Println(fi21, v10, v2)
            v3 = append(v3, (v1 + ("" + ((v1 + v1) + v1))))
            _ = fi21
            _ = v22
        }
        _ = fi17
    }
    var v23 int64 = 9223372036854775807
    for fi24 := 0; fi24 < 4; fi24++ {
        switch v1 {
        case "\n\t":
            fmt.Println(v1, v4, v23)
        }
        if true {
            v5 = v5
            v2 -= v2
        }
        v4 -= v4
        _ = fi24
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        v2 = v10
        v26 := GenS2{F0: "a", F1: (2 <= 65535)}
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        switch (fi27 * fi27) {
        case -277:
            var v28 int16 = v5
            _ = v28
        case -189:
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 int32 = 905
                v5 = 42
                _ = fi29
                _ = v30
            }
        case 32768:
            for fi31 := 0; fi31 < 20; fi31++ {
                v32 := GenS2{F0: ("foo bar" + v1), F1: (!true)}
                var v33 uint64 = 730
                _ = fi31
                _ = v32
                _ = v33
            }
        default:
            v34 := GenS3{F0: true, F1: 1000}
            _ = v34
        }
        var v35 uint64 = 18446744073709551615
        var v36 float32 = 1.5
        _ = fi27
        _ = v35
        _ = v36
    }
    switch 0 {
    case 9223372036854775807:
        _ = len(v3)
    }
    v37 := GenS1{F0: -2147483648, F1: 3}
    v23 -= 0
    for fi38 := 0; fi38 < 7; fi38++ {
        v39 := GenS4{F0: ("foo bar" + (v1 + "hello"))}
        _ = fi38
        _ = v39
    }
    if len(v3) > 0 {
        v3[0] = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
