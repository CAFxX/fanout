package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint32 }
func genFn0(p0 bool) int32 {
    p0 = (!(!(32768 <= -2147483648)))
    v5 := []int{42, -468}
    var v6 int32 = 903
    _ = p0
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 int64, p1 int64, p2 uint32) bool {
    switch "\n\t" {
    case "\n\t":
        p0 += (-(p1))
    case "":
        var v7 uint16 = 2
        _ = v7
    }
    if (745 < 4294967295) {
        switch ("a" + "hello") {
        case "\n\t":
            for fi8 := 0; fi8 < 10; fi8++ {
                p1 = p0
                _ = fi8
            }
        case "foo bar":
            v9 := []int{-2147483648}
            _ = v9
        case "hello":
            switch 0 {
            case 347:
                switch 4294967295 {
                case -2147483648:
                    var v10 float32 = 1.5
                    _ = v10
                case 10:
                    v11 := GenS1{F0: 0, F1: 42}
                    _ = v11
                }
            case 0:
                v12 := GenS2{F0: false}
                _ = v12
            }
        }
    } else {
        var v13 uint64 = 768
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return (true || true)
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]int)
    var v3 uint64 = 42
    var v4 uint = 1
    v14 := GenS1{F0: 822, F1: 0}
    v15 := []int{32768, 539, 0, -2147483648}
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := GenS2{F0: true}
        v18 := GenS3{F0: "foo bar", F1: v4}
        var v19 int8 = 127
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    if true {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 bool = (true && false)
            v15 = append(v15, fi20)
            _ = fi20
            _ = v21
        }
        v22 := GenS2{F0: (true && false)}
        _ = v22
    } else {
        v23 := GenS4{F0: 2147483647}
        fmt.Println(v4, v1)
        _ = v23
    }
    v1 = "foo bar"
    v24 := GenS4{F0: 2}
    v2[65535] = 749
    v15 = append(v15, 299)
    if (0.1 < -2.25) {
        if false {
            for fi25 := 0; fi25 < 16; fi25++ {
                var v26 int64 = 1
                switch (fi25 % (((fi25 & (fi25 | (fi25 + fi25)))) | 1)) {
                case 32767:
                    var v27 string = v1
                    _ = v27
                }
                _ = fi25
                _ = v26
            }
            if len(v15) > 0 {
                v15[0] = 264
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            if ("foo bar" < ((v1 + ((v1 + v1) + ("\n\t" + "hello"))) + (("hello" + (((v1 + v1) + "") + v1)) + v1))) {
                fmt.Println(v4, v4, v4)
                v29 := GenS4{F0: 2}
                fmt.Println(v3)
                _ = v29
            } else {
                var v30 float64 = 1.5
                v15 = append(v15, fi28)
                v4 = v4
                _ = v30
            }
            v4 = 2
            _ = fi28
        }
        if false {
            var v31 int16 = -32768
            fmt.Println(v31, v31, v1)
            switch ((v1 + ("a" + v1)) + v1) {
            case "a":
                _ = len(v2)
            default:
                fmt.Println(v3)
            }
            _ = v31
        } else {
            fmt.Println(v4, v4, v4)
            v32 := GenS2{F0: true}
            if true {
                if (!false) {
                    var v33 uint64 = v3
                    v3 -= v3
                    _ = v33
                } else {
                    for fi34 := 0; fi34 < 5; fi34++ {
                        for fi35 := 0; fi35 < 17; fi35++ {
                            v1 = "\n\t"
                            _ = fi35
                        }
                        var v36 uint64 = v3
                        _ = fi34
                        _ = v36
                    }
                    v37 := v2[134]
                    if (true || false) {
                        var v38 uint16 = 32767
                        var v39 float64 = 1.5
                        _ = v38
                        _ = v39
                    }
                    _ = v37
                }
                v4 = 42
            } else {
                var v40 int32 = 87
                v40 -= int32(v3)
                for fi41 := 0; fi41 < 17; fi41++ {
                    for fi42 := 0; fi42 < 11; fi42++ {
                        fmt.Println(v40)
                        var v43 int = fi41
                        var v44 float32 = 1e10
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    _ = fi41
                }
                _ = v40
            }
            _ = v32
        }
    } else {
        v4 += 4294967295
        v4 -= 42
    }
    fmt.Println(v3, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
