package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int; F1 int64; F2 uint32 }
type GenS4 struct { F0 int }
func genFn0(p0 string) uint16 {
    var v7 uint32 = 2147483647
    if true {
        v8 := []int{-2147483648, 3}
        _ = v8
    } else {
        var v9 uint32 = v7
        v9 -= 2147483647
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS1{F0: 18446744073709551615}
            var v12 bool = (!false)
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    v7 = (^(v7))
    v7 -= v7
    switch p0 {
    case "a":
        switch "foo bar" {
        case "\n\t":
            v13 := GenS1{F0: 42}
            _ = v13
        case "hello":
            v14 := []int{}
            _ = v14
        case "foo bar":
            for fi15 := 0; fi15 < 4; fi15++ {
                p0 = p0
                p0 = p0
                _ = fi15
            }
        }
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 string) bool {
    if true {
        var v16 int16 = 42
        switch (p1 + "hello") {
        case "\n\t":
            if false {
                v17 := []int{-963, 9223372036854775807}
                p0 -= p0
                _ = v17
            } else {
                if (255 > 4294967295) {
                    v16 += 42
                }
            }
        case "":
            if false {
                var v18 uint = 42
                for fi19 := 0; fi19 < 19; fi19++ {
                    v20 := []int{}
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                fmt.Println(p1)
            }
        }
        _ = v16
    } else {
        if (0 > 2) {
            for fi21 := 0; fi21 < 5; fi21++ {
                var v22 uint8 = 2
                _ = fi21
                _ = v22
            }
            v23 := []int{4294967295}
            _ = v23
        }
        var v24 uint = 53
        _ = v24
    }
    if (true && (!(32767 <= -9223372036854775808))) {
        if (0 != 182) {
            var v25 uint = 4294967295
            if true {
                v26 := []int{-193, 469}
                switch -807 {
                case -785:
                    p0 = p0
                }
                fmt.Println(v25)
                _ = v26
            }
            for fi27 := 0; fi27 < 11; fi27++ {
                if false {
                    fmt.Println(p1, v25)
                    if (false || true) {
                        v28 := GenS1{F0: v25}
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 3; fi29++ {
                        if (42 <= 819) {
                            var v30 uint = v25
                            v25 = 18446744073709551615
                            _ = v30
                        }
                        var v31 uint = 18446744073709551615
                        var v32 uint16 = 631
                        _ = fi29
                        _ = v31
                        _ = v32
                    }
                }
                _ = fi27
            }
            _ = v25
        } else {
            var v33 string = "hello"
            v33 = (p1 + v33)
            var v34 uint8 = 255
            _ = v33
            _ = v34
        }
        var v35 uint32 = 65535
        _ = v35
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 int = 690
    v2 := []int{}
    var v3 uint32 = 42
    var v4 uint64 = 42
    var v5 bool = ((934 == 0) && true)
    var v6 uint64 = (v4 / (((v4 & (v4 - v4))) | 1))
    v5 = v5
    v36 := GenS2{F0: 1, F1: (0 == -2147483648)}
    v37 := GenS3{F0: 779, F1: -619, F2: 2147483648}
    var v38 int8 = 42
    for fi39 := 0; fi39 < 14; fi39++ {
        v40 := GenS1{F0: 0}
        _ = fi39
        _ = v40
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        v3 += 905
    case "\n\t":
        var v41 string = "a"
        _ = v41
    }
    var v42 float32 = 1e10
    for fi43 := 0; fi43 < 6; fi43++ {
        v3 -= (^(v3))
        _ = fi43
    }
    v2 = append(v2, -510)
    switch (v1 - int(v6)) {
    case 62:
        for fi44 := 0; fi44 < 20; fi44++ {
            fmt.Println(v42, v4)
            var v45 uint8 = 42
            _ = fi44
            _ = v45
        }
    case 2147483647:
        v3 = v3
    case 782:
        v3 = (v3 * v3)
    }
    v46 := GenS3{F0: v1, F1: 1, F2: v3}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v47 := GenS4{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
