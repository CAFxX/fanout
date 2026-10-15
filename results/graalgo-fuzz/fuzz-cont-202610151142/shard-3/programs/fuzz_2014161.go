package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int; F1 string; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
func genFn0(p0 int16) uint64 {
    for fi5 := 0; fi5 < 2; fi5++ {
        p0 -= p0
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(fi6, p0, fi6)
            _ = fi6
        }
        switch "" {
        case "hello":
            var v7 int8 = -1
            _ = v7
        case "\n\t":
            var v8 int16 = p0
            _ = v8
        case "foo bar":
            v9 := GenS1{F0: -2147483648, F1: (^(fi5))}
            _ = v9
        }
        _ = fi5
    }
    if false {
        for fi10 := 0; fi10 < 16; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                if true {
                    var v12 uint = 53
                    v13 := []int{221, fi10, fi10}
                    _ = v12
                    _ = v13
                } else {
                    switch fi10 {
                    case 128:
                        fmt.Println(p0)
                    case 104:
                        p0 += 435
                    case 32767:
                        v14 := GenS2{F0: 213, F1: "hello", F2: 32767}
                        _ = v14
                    default:
                        switch ("hello" + "a") {
                        case "a":
                            switch (^(fi11)) {
                            case 1:
                                p0 -= p0
                            default:
                                var v15 string = ((("\n\t" + ((("hello" + (("a" + ("\n\t" + ("foo bar" + "foo bar"))) + (((("foo bar" + ("foo bar" + "foo bar")) + "a") + "hello") + (("foo bar" + "") + "a")))) + "") + "\n\t")) + ("" + "\n\t")) + "a")
                                _ = v15
                            }
                        }
                    }
                    switch (fi10 * (fi11 & (fi11 - fi11))) {
                    case 32768:
                        var v16 int = (^(fi11))
                        _ = v16
                    }
                }
                _ = fi11
            }
            v17 := GenS2{F0: fi10, F1: "hello", F2: 42}
            _ = fi10
            _ = v17
        }
        v18 := GenS3{F0: 42, F1: 0, F2: 271}
        _ = v18
    } else {
        var v19 float64 = -2.25
        switch -9223372036854775808 {
        case -345:
            if ((!(258 < 501)) && true) {
                p0 = p0
                var v20 uint64 = 2
                _ = v20
            }
        }
        _ = v19
    }
    _ = p0
    return 42
}
func genFn1(p0 int32, p1 bool) uint64 {
    var v21 bool = (!p1)
    switch ("a" + "hello") {
    case "hello":
        switch 9223372036854775807 {
        case -9223372036854775808:
            var v22 string = ""
            _ = v22
        case 42:
            v23 := GenS2{F0: 32767, F1: ("hello" + ("hello" + ((("a" + (("a" + "\n\t") + "a")) + "") + (("\n\t" + "a") + "\n\t")))), F2: 65535}
            _ = v23
        case -1:
            var v24 string = ("foo bar" + "a")
            _ = v24
        default:
            p1 = false
        }
    case "":
        v25 := []int{}
        _ = v25
    case "a":
        fmt.Println(v21, v21, p0)
    }
    v26 := GenS3{F0: 18446744073709551615, F1: 1, F2: 680}
    switch -697 {
    case 0:
        switch 151 {
        case 1:
            v27 := []int{50, 42}
            _ = v27
        case 961:
            for fi28 := 0; fi28 < 8; fi28++ {
                v29 := []int{fi28, -1, -914}
                var v30 uint32 = 2147483647
                if len(v29) >= 2 {
                    v29 = v29[1:len(v29)]
                }
                _ = fi28
                _ = v29
                _ = v30
            }
        }
    default:
        var v31 int16 = 42
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v26
    return 0
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int8 = 42
    var v3 uint = 980
    var v4 uint = v3
    v32 := GenS4{F0: v1}
    if ("foo bar" != ((("\n\t" + "hello") + (("hello" + "") + ("\n\t" + ("" + "")))) + "foo bar")) {
        if false {
            v33 := GenS3{F0: (v1 & v1), F1: 1, F2: v3}
            v1 = v1
            _ = v33
        }
    } else {
        v34 := []int{-946, -152, 127}
        _ = v34
    }
    v35 := []int{1, 943, 9223372036854775807, 127}
    v1 = 4294967295
    var v36 uint32 = 498
    v4 = (v3 >> (((v4 / ((803) | 1))) & 63))
    var v37 int = 128
    for fi38 := 0; fi38 < 13; fi38++ {
        v39 := GenS5{F0: (fi38 / ((v37) | 1))}
        if false {
            for fi40 := 0; fi40 < 8; fi40++ {
                if len(v35) > 0 {
                    v35[0] = fi40
                }
                _ = fi40
            }
            v41 := GenS2{F0: v37, F1: "hello", F2: 0}
            var v42 int = v37
            _ = v41
            _ = v42
        }
        v37 = v37
        _ = fi38
        _ = v39
    }
    var v43 int16 = 32767
    var v44 int32 = 42
    var v45 string = "a"
    v45 = "a"
    for fi46 := 0; fi46 < 1; fi46++ {
        for fi47 := 0; fi47 < 6; fi47++ {
            var v48 float64 = 1.5
            v3 = v3
            v2 = (v2 + (v2 & v2))
            var v49 bool = (-9223372036854775808 == int(v4))
            _ = fi47
            _ = v48
            _ = v49
        }
        fmt.Println(v2, v4, v45)
        if len(v35) > 0 {
            v35[0] = (v37 + (v37 ^ (fi46 - v37)))
        }
        _ = fi46
    }
    var v50 uint64 = 18446744073709551615
    fmt.Println(v44)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v50
}
