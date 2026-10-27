package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int; F2 string }
type GenS3 struct { F0 uint64; F1 uint32; F2 string }
func genFn0(p0 uint64, p1 int64, p2 int32) uint64 {
    var v6 int64 = -1
    if ((((-(p1)) != (v6 >> ((255) & 63))) && true) || true) {
        v7 := []int{-1, -2147483648, 4294967295}
        p1 = v6
        v8 := GenS1{F0: p2}
        _ = v7
        _ = v8
    }
    if ((!(!(32767 >= 530))) || (float64(0.1) >= -2.25)) {
        v9 := []int{}
        for fi10 := 0; fi10 < 18; fi10++ {
            switch fi10 {
            case -1:
                fmt.Println(p2, v6)
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                fmt.Println(p2)
                _ = fi11
            }
            var v12 string = "hello"
            _ = fi10
            _ = v12
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1() uint {
    var v13 uint = 2
    v14 := GenS2{F0: 962, F1: 32767, F2: "\n\t"}
    fmt.Println(v13)
    _ = v13
    _ = v14
    return 2
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 bool = (!(!true))
    v5 := []string{((("" + "hello") + "foo bar") + "\n\t"), "a", "hello", ("a" + "a")}
    if (1 >= -1) {
        v15 := GenS2{F0: 2, F1: 127, F2: (("a" + (("\n\t" + "foo bar") + "foo bar")) + (("" + "foo bar") + "foo bar"))}
        _ = v15
    }
    if v4 {
        var v16 bool = v4
        _ = v16
    } else {
        if (!v4) {
            var v17 int16 = 795
            for fi18 := 0; fi18 < 7; fi18++ {
                if (913 != 0) {
                    v19 := GenS3{F0: 4294967295, F1: 4294967295, F2: "foo bar"}
                    var v20 uint16 = 65535
                    var v21 uint32 = 1
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    v22 := GenS2{F0: genFn0(128, 0, 42), F1: fi18, F2: ""}
                    fmt.Println(v4, v4, v17)
                    _ = v22
                }
                _ = fi18
            }
            var v23 uint64 = genFn0(2, 0, 0)
            _ = v17
            _ = v23
        }
        var v24 int = 1
        for fi25 := 0; fi25 < 2; fi25++ {
            v26 := GenS3{F0: 18446744073709551615, F1: 2, F2: "a"}
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 uint16 = 32767
        _ = fi27
        _ = v28
    }
    if v4 {
        if (1 >= 29) {
            var v29 uint = 706
            if len(v5) > 0 {
                v5[0] = ("a" + "")
            }
            v30 := GenS2{F0: 944, F1: 128, F2: ""}
            _ = v29
            _ = v30
        }
    }
    if true {
        var v31 int8 = 0
        _ = v31
    }
    var v32 int = 0
    _ = len(v5)
    v4 = (("" + "hello") < "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
