package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 string }
func genFn0() int {
    v6 := []int{32768, 9223372036854775807, 1}
    if (((!(1 != 0)) || ((false || true) || false)) && (118 >= 962)) {
        var v7 uint8 = 0
        var v8 int32 = 1
        v9 := GenS1{F0: 18446744073709551615}
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        fmt.Println(fi10)
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(fi10, fi11)
            if true {
                if (true && (42 != 906)) {
                    var v12 bool = true
                    if v12 {
                        var v13 uint = 0
                        _ = v13
                    } else {
                        var v14 bool = v12
                        _ = v14
                    }
                    if (!((-760 >= 0) && true)) {
                        if ("" != "") {
                            v15 := GenS1{F0: 65535}
                            var v16 int16 = 32767
                            _ = len(v6)
                            _ = v15
                            _ = v16
                        } else {
                            var v17 int8 = 100
                            var v18 uint64 = 914
                            _ = v17
                            _ = v18
                        }
                        var v19 uint64 = 0
                        var v20 int16 = -226
                        _ = v19
                        _ = v20
                    }
                    _ = v12
                }
                var v21 uint = 18446744073709551615
                var v22 uint = (v21 | v21)
                _ = v21
                _ = v22
            } else {
                if (!(!false)) {
                    var v23 int32 = 90
                    _ = v23
                } else {
                    switch (fi10 >> ((928) & 63)) {
                    case 32768:
                        if (42 >= 788) {
                            if false {
                                fmt.Println(fi10, fi11)
                            }
                            v24 := GenS2{F0: 18446744073709551615}
                            var v25 int8 = 0
                            _ = v24
                            _ = v25
                        }
                    case -9223372036854775808:
                        var v26 uint64 = 4294967295
                        _ = v26
                    case -1:
                        var v27 int = fi10
                        _ = v27
                    }
                    var v28 uint = 1
                    if (fi11 != fi10) {
                        var v29 string = "\n\t"
                        var v30 uint8 = 42
                        _ = v29
                        _ = v30
                    }
                    _ = v28
                }
                var v31 uint16 = 0
                _ = v31
            }
            var v32 uint = 2
            _ = fi11
            _ = v32
        }
        v33 := GenS1{F0: 42}
        _ = fi10
        _ = v33
    }
    if false {
        v6 = append(v6, 4294967295)
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
    } else {
        var v34 uint = 2
        _ = v34
    }
    var v35 int = 1
    _ = v35
    _ = v6
    return (v35 + v35)
}

func main() {
    var v1 bool = true
    v2 := []string{("\n\t" + "hello"), ("hello" + ("a" + ""))}
    v3 := make(map[int]uint32)
    v4 := []int{-9223372036854775808}
    var v5 int32 = -1
    var v36 int32 = v5
    v37 := GenS2{F0: 342}
    var v38 bool = v1
    for fi39 := 0; fi39 < 13; fi39++ {
        fmt.Println(v5)
        var v40 bool = v1
        _ = fi39
        _ = v40
    }
    if (!(18446744073709551615 < 85)) {
        v41 := GenS3{F0: 0, F1: 0}
        v42 := GenS4{F0: ("foo bar" + "hello")}
        _ = v41
        _ = v42
    }
    var v43 uint16 = 2
    var v44 uint32 = 956
    v1 = false
    v5 += v36
    fmt.Println(v38, v44, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
