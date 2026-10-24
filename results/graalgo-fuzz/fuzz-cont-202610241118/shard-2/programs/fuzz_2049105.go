package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int8; F2 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "a"
    var v3 int = 1
    if false {
        switch ("" + v2) {
        case "foo bar":
            if (true && (!(18 < 255))) {
                for fi4 := 0; fi4 < 3; fi4++ {
                    fmt.Println(v2)
                    v3 += 42
                    _ = fi4
                }
                v3 += -2147483648
            }
        case "\n\t":
            var v5 string = v2
            _ = v5
        default:
            if (-904 == 1000) {
                var v6 uint = 161
                v2 = ((("\n\t" + v2) + v2) + ("a" + ""))
                _ = v6
            }
        }
        v3 -= -9223372036854775808
    } else {
        switch v2 {
        case "\n\t":
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 int8 = 0
                var v9 uint64 = 2
                _ = fi7
                _ = v8
                _ = v9
            }
        case "hello":
            if false {
                var v10 int32 = 42
                v10 += (v10 - v10)
                _ = v10
            }
        }
        v3 = v3
    }
    if ((0 < 2147483647) && ((54 <= 127) && ((42 > 1) && true))) {
        v11 := GenS1{F0: 100, F1: 127, F2: 0}
        _ = v11
    } else {
        if ((721 == 635) && (!false)) {
            v12 := GenS2{F0: (2 > 255), F1: 545, F2: 2}
            var v13 bool = (false || (496 >= 4294967295))
            switch 127 {
            case 398:
                v14 := []int{}
                _ = v14
            case 127:
                var v15 int8 = 127
                _ = v15
            }
            _ = v12
            _ = v13
        } else {
            v16 := GenS2{F0: (2 <= 1), F1: -1, F2: 255}
            for fi17 := 0; fi17 < 7; fi17++ {
                v3 += v3
                var v18 string = ("hello" + v2)
                v2 = (v18 + (v18 + (("hello" + v18) + "")))
                _ = fi17
                _ = v18
            }
            switch "foo bar" {
            case "":
                v19 := GenS1{F0: 1, F1: 100, F2: 255}
                _ = v19
            default:
                v1[65535] = 2
            }
            _ = v16
        }
        fmt.Println(v2, v2)
        fmt.Println(v2, v3)
    }
    if false {
        fmt.Println(v2, v3)
        var v20 float64 = (3.14159 - 1.5)
        _ = v20
    }
    if (-856 == -2147483648) {
        v3 = (^(v3))
        v1[65535] = 1
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 float32 = (0.1 + -2.25)
            v24 := []int{}
            _ = fi22
            _ = v23
            _ = v24
        }
        v3 += -2147483648
        if false {
            var v25 int8 = -128
            var v26 bool = (1.5 >= 3.14159)
            v27 := GenS3{F0: 255, F1: v25, F2: (!v26)}
            _ = v25
            _ = v26
            _ = v27
        } else {
            v3 += 914
        }
        _ = fi21
    }
    v28 := []int{54}
    switch 547 {
    case -787:
        v29 := GenS1{F0: 0, F1: 86, F2: 0}
        _ = v29
    case 128:
        v30 := GenS2{F0: (-1 != 1000), F1: 0, F2: 128}
        _ = v30
    case -602:
        if true {
            switch (v2 + "") {
            case "":
                v3 -= 128
            }
            _ = len(v28)
        } else {
            if false {
                v2 = "hello"
            }
            switch (v2 + v2) {
            case "a":
                var v31 uint = 42
                _ = v31
            case "\n\t":
                v32 := GenS1{F0: 21, F1: 127, F2: 32767}
                _ = v32
            }
        }
    default:
        if len(v28) >= 2 {
            v28 = v28[1:len(v28)]
        }
    }
    v3 += (v3 & v3)
    for fi33 := 0; fi33 < 10; fi33++ {
        switch 2147483647 {
        case 42:
            v3 -= v3
        default:
            if true {
                if len(v28) > 0 {
                    v28[0] = -1
                }
                v34 := GenS3{F0: 65535, F1: 127, F2: (false || false)}
                _ = v34
            }
        }
        v35 := GenS2{F0: false, F1: 32767, F2: 128}
        _ = fi33
        _ = v35
    }
    fmt.Println(v3, v2)
    var v36 int32 = 42
    v37 := GenS1{F0: -31, F1: 19, F2: 254}
    v3 += v3
    v36 -= v36
    var v38 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
