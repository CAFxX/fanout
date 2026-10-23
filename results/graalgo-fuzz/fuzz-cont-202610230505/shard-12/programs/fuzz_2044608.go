package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int; F1 uint16 }

func main() {
    var v1 float64 = 1.5
    v2 := []string{"\n\t", "foo bar", "\n\t", "hello"}
    var v3 string = ""
    var v4 int64 = -221
    var v5 int16 = -401
    if (true || (v4 > v4)) {
        if (-2147483648 >= 1) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = len(v2)
        }
    } else {
        v3 = (v3 + "foo bar")
    }
    v6 := GenS1{F0: -1, F1: 42, F2: 168}
    v7 := GenS1{F0: 42, F1: 1, F2: 65535}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 string = ""
        _ = fi8
        _ = v9
    }
    var v10 int32 = 0
    v2 = append(v2, "hello")
    if true {
        v5 = int16(v10)
    } else {
        v1 = (v1 + 1e10)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            switch 42 {
            case -892:
                v12 := GenS2{F0: v4}
                _ = v12
            }
            v10 = (v10 / ((v10) | 1))
            _ = fi11
        }
    }
    _ = len(v2)
    if true {
        if true {
            var v13 float64 = 0.1
            _ = v13
        }
        v14 := GenS3{F0: 2, F1: v10}
        v15 := GenS4{F0: 0, F1: 863}
        _ = v14
        _ = v15
    }
    var v16 uint = 0
    if len(v2) > 0 {
        v2[0] = ""
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            switch -2147483648 {
            case 32767:
                var v19 int8 = 3
                _ = v19
            case -523:
                if true {
                    v4 += (v4 + v4)
                }
            default:
                var v20 string = ("" + v3)
                _ = v20
            }
            v1 += (v1 + v1)
            switch 2147483647 {
            case -994:
                var v21 int16 = 0
                _ = v21
            }
            _ = fi18
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 int8 = -1
            var v24 uint64 = 4294967295
            fmt.Println(v5, fi22)
            var v25 uint8 = 2
            var v26 string = v3
            if false {
                var v27 uint64 = v24
                var v28 bool = ((("\n\t" >= "") || (v1 >= 0.1)) || (!false))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v29 := GenS2{F0: (v4 + 527)}
                _ = v27
                _ = v28
                _ = v29
            } else {
                v1 = -2.25
                var v30 int = fi22
                if ((v10 + (v10 - v10)) >= v10) {
                    v23 = (v23 - (v23 & v23))
                }
                _ = v30
            }
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
            _ = v26
        }
        var v31 int = (fi17 - (fi17 + fi17))
        _ = fi17
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
