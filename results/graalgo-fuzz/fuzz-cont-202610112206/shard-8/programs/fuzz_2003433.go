package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 string }
func genFn0(p0 int64) int16 {
    var v4 uint8 = 42
    v5 := GenS1{F0: 4294967295, F1: -924, F2: 32767}
    v6 := GenS1{F0: -48, F1: 999, F2: 65535}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1() bool {
    if false {
        var v7 int64 = -577
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int16 = 42
            if (true && (("foo bar" + "") <= ("hello" + "hello"))) {
                fmt.Println(v7, fi8)
                v10 := []int{(fi8 << ((551) & 63)), fi8}
                for fi11 := 0; fi11 < 17; fi11++ {
                    var v12 int64 = (v7 & 435)
                    if true {
                        var v13 uint32 = 2147483647
                        var v14 string = "hello"
                        v9 -= (v9 ^ (v9 - genFn0((v7 - v12))))
                        _ = v13
                        _ = v14
                    }
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        v15 := GenS2{F0: (!(0 <= 531))}
        _ = v15
        _ = v7
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 string = ("a" + "hello")
            switch ((v17 + v17) + (v17 + "a")) {
            case "":
                v18 := []int{-62}
                _ = v18
            case "a":
                var v19 uint16 = 1
                _ = v19
            case "foo bar":
                var v20 uint16 = 846
                _ = v20
            default:
                for fi21 := 0; fi21 < 4; fi21++ {
                    switch (v17 + (v17 + v17)) {
                    case "hello":
                        v17 = ""
                    case "":
                        v22 := []int{}
                        _ = v22
                    case "a":
                        fmt.Println(v17, fi16, v17)
                    }
                    v23 := []int{fi21, fi21, 128, 32768}
                    var v24 uint32 = 2147483648
                    fmt.Println(fi21, fi16, v17)
                    _ = fi21
                    _ = v23
                    _ = v24
                }
            }
            _ = fi16
            _ = v17
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            fmt.Println(fi25)
            _ = fi25
        }
        var v26 uint16 = 65535
        _ = v26
    }
    var v27 bool = false
    fmt.Println(v27)
    v27 = v27
    if (v27 && (965 < 1)) {
        if ("a" != ("foo bar" + "hello")) {
            var v28 uint32 = 0
            _ = v28
        } else {
            fmt.Println(v27)
            for fi29 := 0; fi29 < 6; fi29++ {
                switch "" {
                case "a":
                    var v30 uint = 2
                    _ = v30
                case "hello":
                    v27 = true
                }
                for fi31 := 0; fi31 < 14; fi31++ {
                    if v27 {
                        v27 = (1 != 372)
                    }
                    _ = fi31
                }
                v27 = v27
                _ = fi29
            }
        }
        for fi32 := 0; fi32 < 9; fi32++ {
            v27 = true
            var v33 bool = false
            v27 = (true || (((float64(1.5) - -2.25) + (1e10 + (0.1 + (3.14159 / (0.1 / float64(0.1)))))) == 1e10))
            _ = fi32
            _ = v33
        }
        var v34 int16 = -1
        _ = v34
    } else {
        var v35 int = 42
        v36 := []int{(v35 ^ v35), v35, v35, v35}
        v37 := GenS1{F0: v35, F1: 4294967295, F2: 1}
        _ = v35
        _ = v36
        _ = v37
    }
    _ = v27
    return true
}

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 32767
    var v3 uint8 = 50
    if true {
        v3 = (v3 & v3)
    } else {
        v38 := GenS3{F0: (v1 << ((944) & 63)), F1: ("foo bar" + ("foo bar" + "\n\t"))}
        switch ("hello" + "a") {
        case "":
            v1 = v1
        }
        _ = v38
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        if true {
            v40 := GenS3{F0: v1, F1: "foo bar"}
            var v41 uint32 = 122
            v2 -= (v2 - 42)
            _ = v40
            _ = v41
        }
        v1 -= v1
        _ = fi39
    }
    fmt.Println(v2, v1, v3)
    var v42 uint64 = 583
    var v43 uint64 = (^(v42))
    v42 = 708
    var v44 int32 = 2147483647
    v45 := GenS3{F0: 432, F1: "hello"}
    fmt.Println(v2, v1, v2)
    v46 := []int{-2147483648, 4294967295}
    if len(v46) > 0 {
        v46[0] = 173
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v46:", len(v46), v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
