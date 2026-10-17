package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint8; F2 int32 }
func genFn0(p0 uint64, p1 bool) int32 {
    var v4 string = ""
    v5 := GenS1{F0: v4}
    var v6 int64 = -750
    switch "a" {
    case "foo bar":
        var v7 int32 = 2147483647
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        p0 -= p0
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int64, p1 uint16) int64 {
    var v9 int8 = -128
    if false {
        p1 = 0
        switch ((("a" + "") + "\n\t") + ("foo bar" + "")) {
        case "foo bar":
            fmt.Println(p0, p0, p1)
        case "hello":
            var v10 float64 = (-2.25 - -2.25)
            _ = v10
        case "":
            p0 -= (p0 ^ -1)
        }
        switch "" {
        case "\n\t":
            v11 := []int{}
            _ = v11
        case "":
            v12 := []int{32767, -454, 1, 42}
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := []int{(fi13 / ((9223372036854775807) | 1)), 128}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v9
    return (^(p0))
}

func main() {
    v1 := []uint32{268, 2147483647}
    var v2 string = ("\n\t" + "\n\t")
    var v3 int32 = 1
    if (false && (1000 == 42)) {
        v2 = (((v2 + ("foo bar" + ("a" + v2))) + v2) + (((v2 + v2) + ("a" + "\n\t")) + "\n\t"))
        v3 += 42
        switch -2147483648 {
        case 0:
            switch 0 {
            case 371:
                var v15 float32 = 0.0
                _ = v15
            }
        case 4294967295:
            if (!false) {
                var v16 int32 = 618
                _ = v16
            } else {
                if false {
                    var v17 uint16 = 65535
                    v1 = append(v1, 4294967295)
                    _ = v17
                }
            }
        case 2147483647:
            for fi18 := 0; fi18 < 5; fi18++ {
                if (57 > 0) {
                    v3 = (-(v3))
                }
                v3 = v3
                _ = fi18
            }
        }
    }
    var v19 int = 1
    v20 := GenS1{F0: "a"}
    if (!false) {
        var v21 int32 = v3
        var v22 int16 = 32767
        _ = v21
        _ = v22
    }
    switch v19 {
    case 414:
        var v23 uint = 2
        _ = v23
    }
    var v24 int16 = -735
    v25 := GenS1{F0: v2}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 bool = true
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi26
            _ = v27
        }
    } else {
        if (42 != 600) {
            v28 := GenS2{F0: 0, F1: 128, F2: (v3 ^ v3)}
            _ = v28
        } else {
            switch v2 {
            case "":
                switch (("hello" + v2) + v2) {
                case "\n\t":
                    for fi29 := 0; fi29 < 3; fi29++ {
                        switch 20 {
                        case -2147483648:
                            fmt.Println(v24)
                        case 128:
                            var v30 bool = (!(!true))
                            _ = v30
                        case 676:
                            var v31 string = v2
                            _ = v31
                        }
                        _ = len(v1)
                        _ = fi29
                    }
                default:
                    for fi32 := 0; fi32 < 5; fi32++ {
                        fmt.Println(v3)
                        v3 -= (-(v3))
                        _ = fi32
                    }
                }
            case "a":
                if len(v1) > 0 {
                    v1[0] = 42
                }
            case "\n\t":
                for fi33 := 0; fi33 < 8; fi33++ {
                    for fi34 := 0; fi34 < 8; fi34++ {
                        var v35 float32 = 1e10
                        v36 := GenS1{F0: ""}
                        v37 := GenS1{F0: "\n\t"}
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    v1 = append(v1, 0)
                    _ = fi33
                }
            }
        }
    }
    v3 = v3
    var v38 int32 = genFn0(0, (316 >= 4294967295))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v38
}
