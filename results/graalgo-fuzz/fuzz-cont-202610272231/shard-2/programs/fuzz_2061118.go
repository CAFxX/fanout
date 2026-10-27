package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint32; F2 int8 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint32 }
func genFn0(p0 string, p1 string) uint {
    if false {
        switch -853 {
        case 32768:
            switch ((p0 + p1) + ((p0 + p0) + ("\n\t" + p0))) {
            case "\n\t":
                if false {
                    var v6 string = ("foo bar" + ((((p1 + ("\n\t" + "")) + p0) + p0) + ((("\n\t" + "hello") + p1) + p0)))
                    if false {
                        v7 := []int{}
                        for fi8 := 0; fi8 < 3; fi8++ {
                            var v9 int16 = 1000
                            _ = fi8
                            _ = v9
                        }
                        _ = v7
                    } else {
                        if (42 == 1) {
                            v10 := []int{-2147483648}
                            var v11 int8 = 0
                            var v12 int16 = -32768
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                    }
                    _ = v6
                }
            }
        case 128:
            var v13 float64 = (0.1 / 1.0)
            _ = v13
        case 127:
            for fi14 := 0; fi14 < 5; fi14++ {
                switch p0 {
                case "a":
                    v15 := []int{(fi14 - fi14), 42, (fi14 | fi14)}
                    _ = v15
                case "hello":
                    for fi16 := 0; fi16 < 9; fi16++ {
                        p0 = p1
                        p1 = p1
                        _ = fi16
                    }
                default:
                    if (!false) {
                        v17 := GenS1{F0: -1}
                        if false {
                            p0 = "foo bar"
                            v18 := []int{}
                            p1 = p1
                            _ = v18
                        }
                        if (1 > 2147483647) {
                            p0 = "\n\t"
                        }
                        _ = v17
                    } else {
                        for fi19 := 0; fi19 < 7; fi19++ {
                            v20 := GenS2{F0: ((p0 + ("foo bar" + ("\n\t" + p0))) + p1), F1: 4294967295, F2: -21}
                            _ = fi19
                            _ = v20
                        }
                        fmt.Println(p0, fi14)
                        for fi21 := 0; fi21 < 2; fi21++ {
                            p1 = ("" + p0)
                            _ = fi21
                        }
                    }
                }
                p1 = (p1 + p0)
                _ = fi14
            }
        default:
            p0 = ("a" + (p0 + "hello"))
        }
    } else {
        p0 = p1
        p1 = "hello"
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := []int{}
        if (42 <= 1) {
            var v24 float32 = (float32(1.5) - -2.25)
            _ = v24
        } else {
            var v25 string = (p0 + "hello")
            v25 = (v25 + "hello")
            if false {
                v26 := GenS1{F0: -506}
                _ = v26
            }
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 int8, p1 int8, p2 uint) uint32 {
    p0 -= (p1 ^ 100)
    v27 := GenS2{F0: "\n\t", F1: 0, F2: (-(p0))}
    var v28 uint8 = 42
    switch 120 {
    case 0:
        for fi29 := 0; fi29 < 2; fi29++ {
            if (-2147483648 >= -402) {
                var v30 int = 0
                p2 += p2
                _ = v30
            }
            v31 := []int{}
            _ = fi29
            _ = v31
        }
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    return 1
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []int{-693, 32767, 456}
    v3 := make(map[string]bool)
    v4 := []string{("\n\t" + "a"), "a"}
    var v5 uint32 = 997
    var v32 int = 1
    v3["\n\t"] = true
    if (false || (!("a" == "\n\t"))) {
        v33 := GenS1{F0: 1}
        v34 := GenS3{F0: v32, F1: 161}
        for fi35 := 0; fi35 < 6; fi35++ {
            switch v32 {
            case -585:
                v36 := GenS4{F0: int32(fi35)}
                _ = v36
            case 485:
                v5 += 2147483648
            case 4294967295:
                if (-633 <= 2147483647) {
                    _ = len(v4)
                }
            }
            for fi37 := 0; fi37 < 4; fi37++ {
                v38 := v3[""]
                switch "\n\t" {
                case "hello":
                    fmt.Println(v5, v5, v32)
                case "\n\t":
                    var v39 uint = 119
                    _ = v39
                default:
                    _ = len(v3)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi37
                _ = v38
            }
            _ = fi35
        }
        _ = v33
        _ = v34
    }
    v40 := GenS5{F0: "hello", F1: genFn1(0, 42, 734)}
    _ = len(v3)
    switch ("foo bar" + "hello") {
    case "\n\t":
        for fi41 := 0; fi41 < 1; fi41++ {
            _ = len(v2)
            var v42 int = v32
            _ = fi41
            _ = v42
        }
    case "":
        switch (v32 / (((v32 - (v32 - v32))) | 1)) {
        case 140:
            v5 = 2147483647
        case 42:
            var v43 bool = false
            _ = v43
        default:
            v5 -= 2
        }
    case "a":
        v44 := GenS1{F0: 32767}
        _ = v44
    default:
        for fi45 := 0; fi45 < 12; fi45++ {
            var v46 uint16 = 255
            _ = fi45
            _ = v46
        }
    }
    var v47 int = v32
    switch "hello" {
    case "foo bar":
        _ = len(v3)
    case "hello":
        v32 += v32
    case "":
        v48 := v3[(("foo bar" + ("" + "\n\t")) + "foo bar")]
        _ = v48
    }
    switch ("a" + "a") {
    case "hello":
        for fi49 := 0; fi49 < 8; fi49++ {
            for fi50 := 0; fi50 < 10; fi50++ {
                var v51 int8 = -33
                if ((v47 + (-(v47))) >= fi49) {
                    v5 = v5
                    if len(v2) > 0 {
                        v2[0] = v47
                    }
                    v47 += fi49
                }
                var v52 uint64 = 983
                _ = fi50
                _ = v51
                _ = v52
            }
            _ = fi49
        }
    case "":
        if len(v2) > 0 {
            v2[0] = -1
        }
    default:
        var v53 uint = 4294967295
        _ = v53
    }
    fmt.Println(v1, v32)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v47
    _ = v5
}
