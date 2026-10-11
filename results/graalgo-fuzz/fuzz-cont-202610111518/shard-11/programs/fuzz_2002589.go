package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 bool) int {
    var v7 uint64 = 632
    var v8 float32 = 0.1
    var v9 int = 81
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (-(v9))
}
func genFn1(p0 uint8, p1 int64, p2 int32) uint {
    v10 := GenS1{F0: 347}
    if true {
        fmt.Println(p1)
        p2 = p2
    }
    var v11 uint = 551
    switch "hello" {
    case "a":
        var v12 int64 = p1
        _ = v12
    case "foo bar":
        if ((!((!true) || false)) || false) {
            p0 = p0
            p1 = (^(p1))
            v11 = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return v11
}

func main() {
    v1 := []uint32{867}
    var v2 int8 = 68
    v3 := []string{(("" + "hello") + ("hello" + "foo bar")), "\n\t"}
    var v4 bool = (true && (65535 == 1))
    var v5 bool = (v4 || false)
    var v6 int8 = v2
    v13 := GenS1{F0: 255}
    var v14 uint8 = 2
    v2 += (-(v6))
    var v15 uint = 2
    switch 800 {
    case 4294967295:
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS1{F0: 65535}
            for fi18 := 0; fi18 < 19; fi18++ {
                v19 := GenS1{F0: 743}
                v1 = append(v1, 65535)
                _ = fi18
                _ = v19
            }
            _ = fi16
            _ = v17
        }
    }
    switch "\n\t" {
    case "":
        fmt.Println(v6, v15, v5)
    case "\n\t":
        _ = len(v3)
    case "hello":
        v4 = true
    }
    v20 := GenS1{F0: 616}
    v4 = (32767 < 185)
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS2{F0: -965}
        _ = fi21
        _ = v22
    }
    var v23 int8 = v6
    if ((42 == 901) && (781 >= 289)) {
        for fi24 := 0; fi24 < 17; fi24++ {
            fmt.Println(v2, v14)
            if false {
                _ = len(v3)
            }
            _ = fi24
        }
        switch "\n\t" {
        case "hello":
            var v25 bool = false
            _ = v25
        case "foo bar":
            if v4 {
                for fi26 := 0; fi26 < 10; fi26++ {
                    if true {
                        var v27 uint8 = (^(v14))
                        v28 := GenS1{F0: 255}
                        switch "foo bar" {
                        case "":
                            v15 -= 4294967295
                        }
                        _ = v27
                        _ = v28
                    } else {
                        switch (fi26 >> ((v15) & 63)) {
                        case 128:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 42:
                            v29 := GenS3{F0: -9223372036854775808}
                            _ = v29
                        default:
                            for fi30 := 0; fi30 < 17; fi30++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v31 string = "foo bar"
                                _ = fi30
                                _ = v31
                            }
                        }
                        var v32 int32 = 1
                        _ = v32
                    }
                    v33 := GenS3{F0: -146}
                    _ = fi26
                    _ = v33
                }
            }
        case "\n\t":
            var v34 uint64 = 42
            _ = v34
        }
    }
    v5 = (true || v5)
    var v35 uint = (v15 << ((v15) & 63))
    fmt.Println(v4)
    v36 := GenS1{F0: 0}
    var v37 uint16 = 42
    v5 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
