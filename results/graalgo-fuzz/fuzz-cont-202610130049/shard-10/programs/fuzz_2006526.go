package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 string }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int; F1 int16 }
func genFn0(p0 string, p1 int8, p2 int32) uint64 {
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            fmt.Println(p2, p0, p0)
            v7 := GenS1{F0: 0, F1: p1, F2: "foo bar"}
            v8 := GenS1{F0: 105, F1: (p1 / ((-1) | 1)), F2: p0}
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    switch 32768 {
    case -2147483648:
        v9 := []int{}
        _ = v9
    }
    if ((!true) && (true || true)) {
        switch p0 {
        case "foo bar":
            v10 := []int{-2147483648, 0, 972, 0}
            _ = v10
        }
        var v11 uint8 = 42
        if false {
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS1{F0: 65535, F1: p1, F2: "hello"}
                _ = fi12
                _ = v13
            }
        } else {
            v14 := []int{}
            _ = v14
        }
        _ = v11
    } else {
        if (0.0 > 0.0) {
            p1 += 100
        }
        var v15 int = 128
        _ = v15
    }
    var v16 int64 = 9223372036854775807
    v17 := GenS1{F0: 4294967295, F1: (p1 ^ (p1 >> ((2) & 63))), F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 int64 = -1
    var v2 uint32 = 2147483647
    v3 := []string{("foo bar" + (((("foo bar" + (("hello" + "\n\t") + ("hello" + "foo bar"))) + "a") + ("\n\t" + (("hello" + "") + "foo bar"))) + "hello")), "\n\t", ("hello" + "foo bar")}
    v4 := []int{}
    v2 = (v2 % (((^(v2))) | 1))
    fmt.Println(v2, v1)
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := GenS1{F0: v2, F1: 127, F2: "a"}
            _ = fi19
            _ = v20
        }
        if ((923 != 0) || true) {
            if false {
                switch 1 {
                case -1:
                    v21 := GenS1{F0: v2, F1: -16, F2: "foo bar"}
                    _ = v21
                }
                v22 := GenS1{F0: (v2 + v2), F1: -1, F2: "a"}
                _ = v22
            }
            v23 := GenS1{F0: 0, F1: 11, F2: ("a" + ("hello" + ""))}
            for fi24 := 0; fi24 < 10; fi24++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v25 int32 = 42
                _ = fi24
                _ = v25
            }
            _ = v23
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi18
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 int32 = 0
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: (v2 * 2147483648), F1: -80, F2: ("a" + "hello")}
    var v29 uint32 = 0
    var v30 bool = true
    switch 9223372036854775807 {
    case 32768:
        for fi31 := 0; fi31 < 16; fi31++ {
            switch (fi31 % (((fi31 + 127)) | 1)) {
            case 32768:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 400:
                for fi32 := 0; fi32 < 14; fi32++ {
                    _ = len(v3)
                    _ = fi32
                }
            }
            _ = fi31
        }
    case 0:
        switch "\n\t" {
        case "a":
            for fi33 := 0; fi33 < 14; fi33++ {
                v34 := GenS1{F0: (^(v2)), F1: 54, F2: "a"}
                v35 := GenS2{F0: "foo bar", F1: "a"}
                _ = fi33
                _ = v34
                _ = v35
            }
        case "hello":
            v36 := GenS2{F0: ((("\n\t" + "hello") + ("foo bar" + "a")) + "\n\t"), F1: "hello"}
            _ = v36
        }
    }
    switch (((("hello" + (((("\n\t" + ((("" + ("foo bar" + "a")) + (("hello" + "") + "foo bar")) + "a")) + "a") + "a") + "foo bar")) + "hello") + ("foo bar" + "hello")) + "\n\t") {
    case "\n\t":
        v37 := GenS2{F0: "", F1: ("" + "a")}
        _ = v37
    case "a":
        for fi38 := 0; fi38 < 7; fi38++ {
            if (93 > 97) {
                var v39 int16 = 0
                fmt.Println(v39)
                _ = v39
            } else {
                v40 := GenS3{F0: fi38, F1: -1}
                _ = v40
            }
            switch "hello" {
            case "\n\t":
                if v30 {
                    fmt.Println(fi38)
                }
            }
            _ = fi38
        }
    case "":
        var v41 int64 = v1
        _ = v41
    }
    v29 = (v29 * (v2 * (v29 / ((v29) | 1))))
    v3 = append(v3, "\n\t")
    for fi42 := 0; fi42 < 11; fi42++ {
        fmt.Println(v30, v2)
        if (9223372036854775807 > (v1 % ((-9223372036854775808) | 1))) {
            fmt.Println(v30)
            fmt.Println(v29)
        } else {
            fmt.Println(fi42, v29, v2)
        }
        _ = fi42
    }
    v30 = v30
    var v43 uint = 255
    v2 = 0
    v44 := GenS1{F0: (v29 - uint32(v43)), F1: 0, F2: "foo bar"}
    var v45 float32 = (float32(((float64(1.5) - (0.1 / (3.14159 - float64((0.1 + 1.5))))) * ((1.5 - 0.1) / (float64(((0.0 / (((1.5 - float32(0.1)) * 3.14159) / float32(1e10))) - -2.25)) - 1e10)))) + (((0.1 * (float32((-2.25 / 1.5)) * 1.5)) - 3.14159) - 1e10))
    var v46 int64 = (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
